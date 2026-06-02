package internalSql

import (
	"context"
	"database/sql"
	"fmt"
)

type txKey struct{}

type Transactor interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type sqlTransactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) Transactor {
	return &sqlTransactor{db: db}
}

func (t *sqlTransactor) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	ctxTx := context.WithValue(ctx, txKey{}, tx)

	if err := fn(ctxTx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func Executor(ctx context.Context, db DBTX) DBTX {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return db
}
