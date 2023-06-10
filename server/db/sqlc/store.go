package db

import (
	"database/sql"
)


type Store interface {
	Querier
}
type SQLStore struct { //will talk to real database
	db *sql.DB
	*Queries
}

// NewStore creates a new store interface with SQL store struct
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),// a new query object
	}
}
