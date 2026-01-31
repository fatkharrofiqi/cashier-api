package service

import (
	"cashier-api/internal/model"
	"cashier-api/internal/repository"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	GetByID(id int) (*model.Product, error)
	Create(product model.Product) (*model.Product, error)
	Update(id int, product model.Product) (*model.Product, error)
	Delete(id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id int) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(product model.Product) (*model.Product, error) {
	return s.repo.Create(product)
}

func (s *productService) Update(id int, product model.Product) (*model.Product, error) {
	return s.repo.Update(id, product)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}
