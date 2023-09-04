package util

import (
	"errors"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

var (
	ErrDisconnected = errors.New("disconnected from rabbitmq, trying to reconnect")
)

const (
	// When reconnecting to the server after connection failure
	reconnectDelay = 5 * time.Second

	// When resending messages the server didn't confirm
	resendDelay = 5 * time.Second
)

type Client struct {
	pushQueue     string
	logger        zerolog.Logger
	connection    *amqp.Connection
	channel       *amqp.Channel
	done          chan os.Signal
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
	isConnected   bool
	alive         bool
}

type RabbitMqClient interface {
	handleReconnect(string)
	connect(string) bool
	changeConnection(*amqp.Connection, *amqp.Channel)
	Push([]byte) error
	UnsafePush([]byte) error
	Close() error
}

// New is a constructor that takes address, push and listen queue names, logger, and a channel that will notify rabbitmq client on server shutdown. We calculate the number of threads, create the client, and start the connection process. Connect method connects to the rabbitmq server and creates push/listen channels if they don't exist.
func New(pushQueue, addr string, l zerolog.Logger, done chan os.Signal) *Client {
	client := Client{
		logger:    l,
		pushQueue: pushQueue,
		done:      done,
		alive:     true,
	}

	go client.handleReconnect(addr)
	return &client
}

// handleReconnect will wait for a connection error on
// notifyClose, and then continuously attempt to reconnect.
func (c *Client) handleReconnect(addr string) {
	for c.alive {
		c.isConnected = false
		t := time.Now()
		fmt.Printf("Attempting to connect to rabbitMQ: %s\n", addr)
		var retryCount int
		for !c.connect(addr) {
			if !c.alive {
				return
			}
			select {
			case <-c.done:
				return
			case <-time.After(reconnectDelay + time.Duration(retryCount)*time.Second):
				c.logger.Printf("disconnected from rabbitMQ and failed to connect")
				retryCount++
			}
		}
		c.logger.Printf("Connected to rabbitMQ in: %vms", time.Since(t).Milliseconds())
		select {
		case <-c.done:
			return
		case <-c.notifyClose:
		}
	}
}

// connect will make a single attempt to connect to
// RabbitMq. It returns the success of the attempt.
func (c *Client) connect(addr string) bool {
	conn, err := amqp.Dial(addr)
	if err != nil {
		c.logger.Printf("failed to dial rabbitMQ server: %v", err)
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		c.logger.Printf("failed connecting to channel: %v", err)
		return false
	}
	ch.Confirm(false)

	_, err = ch.QueueDeclare(
		c.pushQueue,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		c.logger.Printf("failed to declare push queue: %v", err)
		return false
	}
	c.changeConnection(conn, ch)
	c.isConnected = true
	return true
}

// changeConnection takes a new connection to the queue,
// and updates the channel listeners to reflect this.
func (c *Client) changeConnection(connection *amqp.Connection, channel *amqp.Channel) {
	c.connection = connection
	c.channel = channel
	c.notifyClose = make(chan *amqp.Error)
	c.notifyConfirm = make(chan amqp.Confirmation)
	c.channel.NotifyClose(c.notifyClose)
	c.channel.NotifyPublish(c.notifyConfirm)
}

// Push will push data onto the queue, and wait for a confirmation.
// If no confirms are received until within the resendTimeout,
// it continuously resends messages until a confirmation is received.
// This will block until the server sends a confirm.
func (c *Client) Push(data []byte) error {
	if !c.isConnected {
		return errors.New("failed to push push: not connected")
	}
	for {
		err := c.UnsafePush(data)
		if err != nil {
			if err == ErrDisconnected {
				continue
			}
			return err
		}
		select {
		case confirm := <-c.notifyConfirm:
			if confirm.Ack {
				return nil
			}
		case <-time.After(resendDelay):
		}
	}
}

// UnsafePush will push to the queue without checking for
// confirmation. It returns an error if it fails to connect.
// No guarantees are provided for whether the server will
// receive the message.
func (c *Client) UnsafePush(data []byte) error {
	if !c.isConnected {
		return ErrDisconnected
	}
	return c.channel.Publish(
		"",          // Exchange
		c.pushQueue, // Routing key
		false,       // Mandatory
		false,       // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}

func (c *Client) Close() error {
	if !c.isConnected {
		return nil
	}
	c.alive = false
	fmt.Println("Waiting for current messages to be processed...")

	err := c.channel.Close()
	if err != nil {
		return err
	}
	err = c.connection.Close()
	if err != nil {
		return err
	}

	c.isConnected = false
	fmt.Println("gracefully stopped rabbitMQ connection")
	return nil
}
