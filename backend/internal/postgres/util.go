package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// withTransaction handles the transaction lifecycle for database access objects.
func withTransaction(ctx context.Context, dbx *sqlx.DB, fn func(*sqlx.Tx) error) error {
	tx, err := dbx.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Println("failed to rollback transaction", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(); err != nil {
				log.Println("failed to commit transaction", commitErr)
			}
		}
	}()

	err = fn(tx)

	return err
}
