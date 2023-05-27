package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/thewackyindian/3iOj/api"
	db "github.com/thewackyindian/3iOj/db/sqlc"
)

const (
	dbDriver = "postgres"
	dBSource = "postgresql://root:secret@localhost:5432/3iOj?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server :" , err)
	}
}