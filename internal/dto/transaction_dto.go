package dto

import "time"

type CheckoutRequest struct {
	Items []CheckoutItemRequest `json:"items"`
}

type CheckoutItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutResponse struct {
	TransactionID int                  `json:"transaction_id"`
	TotalAmount   int                  `json:"total_amount"`
	CreatedAt     time.Time            `json:"created_at"`
	Items         []CheckoutItemDetail `json:"items"`
}

type CheckoutItemDetail struct {
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Subtotal  int    `json:"subtotal"`
}
