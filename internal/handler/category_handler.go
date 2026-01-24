package handler

import (
	"cashier-api/internal/model"
	"cashier-api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) Handlecategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.service.GetAll())
	case http.MethodPost:
		var c model.Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(h.service.Create(c))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if c, ok := h.service.GetByID(id); ok {
			json.NewEncoder(w).Encode(c)
			return
		}
	case http.MethodPut:
		var c model.Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if updated, ok := h.service.Update(id, c); ok {
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

	http.Error(w, "Category not found", http.StatusNotFound)
}
