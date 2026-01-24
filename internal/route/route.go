package route

import (
	"cashier-api/internal/handler"
	"net/http"
)

type Route struct {
	productHandler  *handler.ProductHandler
	categoryHandler *handler.CategoryHandler
}

func NewRoute(productHandler *handler.ProductHandler, categoryHandler *handler.CategoryHandler) *Route {
	return &Route{
		productHandler:  productHandler,
		categoryHandler: categoryHandler,
	}
}

func (r *Route) Register() {
	// Product routes
	http.HandleFunc("/api/product", r.productHandler.HandleProducts)
	http.HandleFunc("/api/product/", r.productHandler.HandleProductByID)

	// Category routes
	http.HandleFunc("/api/category", r.categoryHandler.Handlecategory)
	http.HandleFunc("/api/category/", r.categoryHandler.HandleCategoryByID)

	// Health check
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", healthHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","message":"API running"}`))
}
