package dto

type ReportResponse struct {
	TotalRevenue       int                    `json:"total_revenue"`
	TotalTransactions  int                    `json:"total_transactions"`
	BestSellingProduct *BestSellingProductDTO `json:"best_selling_product,omitempty"`
}

type BestSellingProductDTO struct {
	Name         string `json:"name"`
	QuantitySold int    `json:"quantity_sold"`
}
