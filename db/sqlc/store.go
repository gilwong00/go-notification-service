package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	ExecTx(ctx context.Context, fn func(*Queries) error) error
}

type DBStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &DBStore{
		db:      db,
		Queries: New(db),
	}
}

func (s *DBStore) ExecTx(
	ctx context.Context,
	fn func(*Queries) error,
) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := New(tx)
	if err = fn(query); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
