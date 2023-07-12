package db

import (
	"context"
	"database/sql"
	"fmt"
)

// this interface has all queries and transactions
type Store interface {
	Querier
	execTx(ctx context.Context, fn func(*Queries) error) error
	// UpdateContestTx(ctx context.Context, arg UpdateContestTxParams) (UpdateContestTxResponse, error)
}
type SQLStore struct { //will talk to real database
	db *sql.DB
	*Queries
}

// NewStore creates a new store interface with SQL store struct
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db), // a new query object
	}
}
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	//for executing  a generic DB transaction
	/*
		The idea is simple: this function takes a context and a callback function as input, then it will
		start a new db transaction, create a new Queries object with that transaction, call the callback
		function with the created Queries, and finally commit or rollback the transaction based on the error
		returned by that function.
	*/
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
