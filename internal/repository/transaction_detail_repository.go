package repository

import (
	"cashier-api/internal/model"
	"cashier-api/internal/uow"
	"context"
	"database/sql"
	"fmt"
)

type TransactionDetailRepository interface {
	Create(ctx context.Context, detail model.TransactionDetail) (*model.TransactionDetail, error)
}

type transactionDetailRepository struct {
	db *sql.DB
}

func NewTransactionDetailRepository(db *sql.DB) TransactionDetailRepository {
	return &transactionDetailRepository{db: db}
}

func (r *transactionDetailRepository) Create(ctx context.Context, detail model.TransactionDetail) (*model.TransactionDetail, error) {
	executor := uow.GetExecutor(ctx, r.db)
	query := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id"
	err := executor.QueryRowContext(ctx, query, detail.TransactionID, detail.ProductID, detail.Quantity, detail.Subtotal).Scan(&detail.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction detail: %w", err)
	}
	return &detail, nil
}
