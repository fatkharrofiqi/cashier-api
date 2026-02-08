package repository

import (
	"cashier-api/internal/model"
	"cashier-api/internal/uow"
	"context"
	"database/sql"
	"fmt"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction model.Transaction) (*model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
	executor := uow.GetExecutor(ctx, r.db)
	query := "INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id, created_at"
	err := executor.QueryRowContext(ctx, query, transaction.TotalAmount).Scan(&transaction.ID, &transaction.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
	return &transaction, nil
}
