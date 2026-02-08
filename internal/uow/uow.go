package uow

import (
	"context"
	"database/sql"
)

// TxKey is the context key for storing the active transaction.
type ctxKey string

const TxKey ctxKey = "tx"

// UnitOfWork encapsulates transaction boundaries and exposes repositories
// bound to active transaction.
type UnitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

// Do runs fn within a transaction. It begins a transaction, injects it
// into the context so repositories pick it up via getExecutor, and commits
// or rolls back based on fn's result.
func (u *UnitOfWork) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	ctxWithTx := context.WithValue(ctx, TxKey, tx)

	if err := fn(ctxWithTx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// getExecutor returns a transaction if one exists in the context,
// otherwise returns the database connection.
func GetExecutor(ctx context.Context, db *sql.DB) Executor {
	if tx, ok := ctx.Value(TxKey).(*sql.Tx); ok {
		return tx
	}
	return db
}

// Executor is an interface that both *sql.DB and *sql.Tx implement
// for query operations.
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
