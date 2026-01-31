package handler

import (
	"cashier-api/internal/dto"
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
		products, err := h.service.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responses := make([]dto.ProductResponse, len(products))
		for i, p := range products {
			responses[i] = h.toProductResponse(p)
		}
		json.NewEncoder(w).Encode(responses)
	case http.MethodPost:
		var req dto.ProductCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		product, err := h.service.Create(model.Product{
			Name:       req.Name,
			Price:      req.Price,
			Stock:      req.Stock,
			CategoryID: req.CategoryID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(h.toProductResponse(*product))
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
		p, err := h.service.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(h.toProductResponse(*p))
	case http.MethodPut:
		var req dto.ProductUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		updated, err := h.service.Update(id, model.Product{
			Name:       req.Name,
			Price:      req.Price,
			Stock:      req.Stock,
			CategoryID: req.CategoryID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(h.toProductResponse(*updated))
	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Delete success",
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) toProductResponse(p model.Product) dto.ProductResponse {
	var category *dto.CategoryDTO
	if p.Category != nil {
		category = &dto.CategoryDTO{
			ID:          p.Category.ID,
			Name:        p.Category.Name,
			Description: p.Category.Description,
		}
	}
	return dto.ProductResponse{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
		Category:   category,
	}
}

