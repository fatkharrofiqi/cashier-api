package repository

import (
	"cashier-api/internal/model"
	"cashier-api/internal/uow"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type TransactionDetailRepository interface {
	Create(ctx context.Context, detail model.TransactionDetail) (*model.TransactionDetail, error)
	BulkCreate(ctx context.Context, details []model.TransactionDetail) ([]model.TransactionDetail, error)
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

func (r *transactionDetailRepository) BulkCreate(ctx context.Context, details []model.TransactionDetail) ([]model.TransactionDetail, error) {
	if len(details) == 0 {
		return details, nil
	}

	executor := uow.GetExecutor(ctx, r.db)

	placeholders := make([]string, 0, len(details)*4)
	args := make([]interface{}, 0, len(details)*4)

	for i, detail := range details {
		base := i * 4
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", base+1, base+2, base+3, base+4))
		args = append(args, detail.TransactionID, detail.ProductID, detail.Quantity, detail.Subtotal)
	}

	query := fmt.Sprintf("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES %s RETURNING id, transaction_id, product_id, quantity, subtotal", strings.Join(placeholders, ", "))

	rows, err := executor.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk create transaction details: %w", err)
	}
	defer rows.Close()

	result := make([]model.TransactionDetail, 0, len(details))
	for rows.Next() {
		var detail model.TransactionDetail
		if err := rows.Scan(&detail.ID, &detail.TransactionID, &detail.ProductID, &detail.Quantity, &detail.Subtotal); err != nil {
			return nil, fmt.Errorf("failed to scan transaction detail: %w", err)
		}
		result = append(result, detail)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transaction details: %w", err)
	}

	return result, nil
}
