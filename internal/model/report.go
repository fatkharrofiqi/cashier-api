package model

type Report struct {
	TotalRevenue      int `db:"total_revenue"`
	TotalTransactions int `db:"total_transactions"`
}

type BestSellingProduct struct {
	Name         string `db:"name"`
	QuantitySold  int    `db:"quantity_sold"`
}
