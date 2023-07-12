package token

import "time"

// General token maker interface to manage the
// creation and verification of tokens
// then we have separate JWT and Paseto structs that implement this interface

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
