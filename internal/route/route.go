package route

import (
	"cashier-api/internal/handler"
	"net/http"
)

type Route struct {
	productHandler *handler.ProductHandler
}

func NewRoute(productHandler *handler.ProductHandler) *Route {
	return &Route{
		productHandler: productHandler,
	}
}

func (r *Route) Register() {
	// Product routes
	http.HandleFunc("/api/product", r.productHandler.HandleProducts)
	http.HandleFunc("/api/product/", r.productHandler.HandleProductByID)

	// Health check
	http.HandleFunc("/health", healthHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","message":"API running"}`))
}
