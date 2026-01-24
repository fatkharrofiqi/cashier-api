package repository

import "cashier-api/internal/model"

type ProductRepository interface {
	FindAll() []model.Product
	FindByID(id int) (*model.Product, bool)
	Create(product model.Product) model.Product
	Update(id int, product model.Product) (*model.Product, bool)
	Delete(id int) bool
}

type productRepository struct {
	products []model.Product
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		products: []model.Product{
			{ID: 1, Name: "Indomie Goreng", Price: 3500, Stock: 10},
			{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
		},
	}
}

func (r *productRepository) FindAll() []model.Product {
	return r.products
}

func (r *productRepository) FindByID(id int) (*model.Product, bool) {
	for _, p := range r.products {
		if p.ID == id {
			return &p, true
		}
	}
	return nil, false
}

func (r *productRepository) Create(product model.Product) model.Product {
	product.ID = len(r.products) + 1
	r.products = append(r.products, product)
	return product
}

func (r *productRepository) Update(id int, product model.Product) (*model.Product, bool) {
	for i := range r.products {
		if r.products[i].ID == id {
			product.ID = id
			r.products[i] = product
			return &product, true
		}
	}
	return nil, false
}

func (r *productRepository) Delete(id int) bool {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return true
		}
	}
	return false
}
