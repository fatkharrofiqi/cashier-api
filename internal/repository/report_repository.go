package repository

import (
	"cashier-api/internal/model"
	"database/sql"
	"fmt"
)

type ReportRepository interface {
	GetTodayReport() (*model.Report, *model.BestSellingProduct, error)
	GetReportByDateRange(startDate, endDate string) (*model.Report, *model.BestSellingProduct, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetTodayReport() (*model.Report, *model.BestSellingProduct, error) {
	query := `
		SELECT
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transactions
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`
	var report model.Report
	err := r.db.QueryRow(query).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, fmt.Errorf("failed to get report summary: %w", err)
	}

	query = `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as quantity_sold
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.id, p.name
		ORDER BY quantity_sold DESC
		LIMIT 1
	`
	var bestSellingProduct model.BestSellingProduct
	err = r.db.QueryRow(query).Scan(&bestSellingProduct.Name, &bestSellingProduct.QuantitySold)
	if err == sql.ErrNoRows {
		return &report, nil, nil
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get best selling product: %w", err)
	}

	return &report, &bestSellingProduct, nil
}

func (r *reportRepository) GetReportByDateRange(startDate, endDate string) (*model.Report, *model.BestSellingProduct, error) {
	return r.getReport(startDate, endDate)
}

func (r *reportRepository) getReport(startDate, endDate string) (*model.Report, *model.BestSellingProduct, error) {
	query := `
		SELECT
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transactions
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1::date AND $2::date
	`
	var report model.Report
	err := r.db.QueryRow(query, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, fmt.Errorf("failed to get report summary: %w", err)
	}

	query = `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as quantity_sold
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) BETWEEN $1::date AND $2::date
		GROUP BY p.id, p.name
		ORDER BY quantity_sold DESC
		LIMIT 1
	`
	var bestSellingProduct model.BestSellingProduct
	err = r.db.QueryRow(query, startDate, endDate).Scan(&bestSellingProduct.Name, &bestSellingProduct.QuantitySold)
	if err == sql.ErrNoRows {
		return &report, nil, nil
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get best selling product: %w", err)
	}

	return &report, &bestSellingProduct, nil
}
