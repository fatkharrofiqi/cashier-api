package service

import (
	"cashier-api/internal/dto"
	"cashier-api/internal/model"
	"cashier-api/internal/repository"
)

type ReportService interface {
	GetTodayReport() (*dto.ReportResponse, error)
	GetReportByDateRange(startDate, endDate string) (*dto.ReportResponse, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GetTodayReport() (*dto.ReportResponse, error) {
	report, bestSellingProduct, err := s.repo.GetTodayReport()
	if err != nil {
		return nil, err
	}

	return s.buildReportResponse(report, bestSellingProduct), nil
}

func (s *reportService) GetReportByDateRange(startDate, endDate string) (*dto.ReportResponse, error) {
	report, bestSellingProduct, err := s.repo.GetReportByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return s.buildReportResponse(report, bestSellingProduct), nil
}

func (s *reportService) buildReportResponse(report *model.Report, bestSellingProduct *model.BestSellingProduct) *dto.ReportResponse {
	response := &dto.ReportResponse{
		TotalRevenue:      report.TotalRevenue,
		TotalTransactions: report.TotalTransactions,
	}

	if bestSellingProduct != nil {
		response.BestSellingProduct = &dto.BestSellingProductDTO{
			Name:         bestSellingProduct.Name,
			QuantitySold: bestSellingProduct.QuantitySold,
		}
	}

	return response
}
