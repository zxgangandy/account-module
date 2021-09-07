package dao

import (
	"context"
	"database/sql"
)

type ITransaction interface {
	Transaction(ctx context.Context, fn func(txContext context.Context) error) error
}

func ExecSqlWithTransaction(db *sql.DB, handle func(tx *sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err = handle(tx); err != nil {
		return err
	}
	return tx.Commit()
}
