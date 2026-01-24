package service

import (
	"cashier-api/internal/model"
	"cashier-api/internal/repository"
)

type ProductService interface {
	GetAll() []model.Product
	GetByID(id int) (*model.Product, bool)
	Create(product model.Product) model.Product
	Update(id int, product model.Product) (*model.Product, bool)
	Delete(id int) bool
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAll() []model.Product {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id int) (*model.Product, bool) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(product model.Product) model.Product {
	return s.repo.Create(product)
}

func (s *productService) Update(id int, product model.Product) (*model.Product, bool) {
	return s.repo.Update(id, product)
}

func (s *productService) Delete(id int) bool {
	return s.repo.Delete(id)
}
