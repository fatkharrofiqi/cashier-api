package model

import "time"

type Transaction struct {
	ID         int       `db:"id"`
	TotalAmount int      `db:"total_amount"`
	CreatedAt  time.Time `db:"created_at"`
}

type TransactionDetail struct {
	ID            int `db:"id"`
	TransactionID int `db:"transaction_id"`
	ProductID     int `db:"product_id"`
	Quantity      int `db:"quantity"`
	Subtotal      int `db:"subtotal"`
}
