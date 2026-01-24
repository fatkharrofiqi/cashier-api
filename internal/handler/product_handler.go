package handler

import (
	"cashier-api/internal/model"
	"cashier-api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.service.GetAll())
	case http.MethodPost:
		var p model.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(h.service.Create(p))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if p, ok := h.service.GetByID(id); ok {
			json.NewEncoder(w).Encode(p)
			return
		}
	case http.MethodPut:
		var p model.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if updated, ok := h.service.Update(id, p); ok {
			json.NewEncoder(w).Encode(updated)
			return
		}
	case http.MethodDelete:
		if h.service.Delete(id) {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Delete success",
			})
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
