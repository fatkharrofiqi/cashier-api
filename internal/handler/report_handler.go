package handler

import (
	"cashier-api/internal/service"
	"encoding/json"
	"net/http"
	"time"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(service service.ReportService) *ReportHandler {
	return &ReportHandler{service}
}

func (h *ReportHandler) HandleTodayReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response, err := h.service.GetTodayReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h *ReportHandler) HandleDateRangeReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	if err := validateDate(startDate); err != nil {
		http.Error(w, "Invalid start_date format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateDate(endDate); err != nil {
		http.Error(w, "Invalid end_date format: "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.service.GetReportByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func validateDate(dateStr string) error {
	_, err := time.Parse("2006-01-02", dateStr)
	return err
}
