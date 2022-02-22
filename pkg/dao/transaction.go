package dao

import (
	"context"
	"database/sql"
)

type ITransaction interface {
	Transaction(ctx context.Context, fn func(txContext context.Context) error) error
}

func ExecTx(db *sql.DB, handle func(tx *sql.Tx) error) (err error) {
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

func ExecTxResult(db *sql.DB, handle func(tx *sql.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	value, err := handle(tx)
	if err != nil {
		return nil, err
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return nil, commitErr
	}

	return value, nil
}
