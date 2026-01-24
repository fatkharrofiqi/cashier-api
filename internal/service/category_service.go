package service

import (
	"cashier-api/internal/model"
	"cashier-api/internal/repository"
)

type CategoryService interface {
	GetAll() []model.Category
	GetByID(id int) (*model.Category, bool)
	Create(category model.Category) model.Category
	Update(id int, category model.Category) (*model.Category, bool)
	Delete(id int) bool
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetAll() []model.Category {
	return s.repo.FindAll()
}

func (s *categoryService) GetByID(id int) (*model.Category, bool) {
	return s.repo.FindByID(id)
}

func (s *categoryService) Create(category model.Category) model.Category {
	return s.repo.Create(category)
}

func (s *categoryService) Update(id int, category model.Category) (*model.Category, bool) {
	return s.repo.Update(id, category)
}

func (s *categoryService) Delete(id int) bool {
	return s.repo.Delete(id)
}
