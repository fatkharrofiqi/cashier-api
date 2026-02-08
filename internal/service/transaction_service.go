package service

import (
	"cashier-api/internal/dto"
	"cashier-api/internal/model"
	"cashier-api/internal/repository"
	"cashier-api/internal/uow"
	"context"
	"fmt"
)

type TransactionService interface {
	Checkout(req dto.CheckoutRequest) (*dto.CheckoutResponse, error)
}

type transactionService struct {
	productRepo           repository.ProductRepository
	transactionRepo       repository.TransactionRepository
	transactionDetailRepo repository.TransactionDetailRepository
	uow                   *uow.UnitOfWork
}

func NewTransactionService(
	productRepo repository.ProductRepository,
	transactionRepo repository.TransactionRepository,
	transactionDetailRepo repository.TransactionDetailRepository,
	uow *uow.UnitOfWork,
) TransactionService {
	return &transactionService{
		productRepo:           productRepo,
		transactionRepo:       transactionRepo,
		transactionDetailRepo: transactionDetailRepo,
		uow:                   uow,
	}
}

func (s *transactionService) Checkout(req dto.CheckoutRequest) (*dto.CheckoutResponse, error) {
	if len(req.Items) == 0 {
		return nil, fmt.Errorf("no items in checkout request")
	}

	var result *dto.CheckoutResponse

	err := s.uow.Do(context.Background(), func(ctx context.Context) error {
		items := make([]dto.CheckoutItemDetail, 0, len(req.Items))
		totalAmount := 0

		productData := make(map[int]*model.Product)

		for _, item := range req.Items {
			product, err := s.productRepo.FindByIDForUpdate(ctx, item.ProductID)
			if err != nil {
				return fmt.Errorf("product with id %d not found: %w", item.ProductID, err)
			}

			if product.Stock < item.Quantity {
				return fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)", product.Name, product.Stock, item.Quantity)
			}

			subtotal := product.Price * item.Quantity
			totalAmount += subtotal

			items = append(items, dto.CheckoutItemDetail{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Subtotal:  subtotal,
			})

			productData[item.ProductID] = product
		}

		transaction := model.Transaction{
			TotalAmount: totalAmount,
		}

		createdTransaction, err := s.transactionRepo.Create(ctx, transaction)
		if err != nil {
			return fmt.Errorf("failed to create transaction: %w", err)
		}

		details := make([]model.TransactionDetail, 0, len(items))
		for _, item := range items {
			details = append(details, model.TransactionDetail{
				TransactionID: createdTransaction.ID,
				ProductID:     item.ProductID,
				Quantity:      item.Quantity,
				Subtotal:      item.Subtotal,
			})
		}

		_, err = s.transactionDetailRepo.BulkCreate(ctx, details)
		if err != nil {
			return fmt.Errorf("failed to create transaction details: %w", err)
		}

		for _, item := range req.Items {
			if err := s.productRepo.UpdateStock(ctx, item.ProductID, item.Quantity); err != nil {
				return fmt.Errorf("failed to update stock for product %d: %w", item.ProductID, err)
			}
		}

		result = &dto.CheckoutResponse{
			TransactionID: createdTransaction.ID,
			TotalAmount:   createdTransaction.TotalAmount,
			CreatedAt:     createdTransaction.CreatedAt,
			Items:         items,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
