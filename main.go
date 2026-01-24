package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var product = []Product{
	{ID: 1, Name: "Indomie Goreng", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range product {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product not available yet", http.StatusNotFound)
}

func updateProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range product {
		if product[i].ID == id {
			product[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not available yet", http.StatusNotFound)
}

func deleteProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for i, p := range product {
		if p.ID == id {
			product = append(product[:i], product[1+i:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Delete success",
			})
			return
		}
	}

	http.Error(w, "Product not available yet", http.StatusNotFound)
}

func main() {
	// GET localhost:8080/api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductByID(w, r)
		case "PUT":
			updateProductByID(w, r)
		case "DELETE":
			deleteProductByID(w, r)
		}
	})

	// GET localhost:8080/api/product
	// POST localhost:8080/api/product
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(product)
		case "POST":
			var productBaru Product
			err := json.NewDecoder(r.Body).Decode(&productBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			productBaru.ID = len(product) + 1
			product = append(product, productBaru)

			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(productBaru)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "Ok",
			"message": "Api Running",
		})
	})

	fmt.Println("Server running on localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed starting server")
	}
}
