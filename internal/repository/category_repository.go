package repository

import "cashier-api/internal/model"

type CategoryRepository interface {
	FindAll() []model.Category
	FindByID(id int) (*model.Category, bool)
	Create(category model.Category) model.Category
	Update(id int, category model.Category) (*model.Category, bool)
	Delete(id int) bool
}

type categoryRepository struct {
	category []model.Category
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		category: []model.Category{
			{ID: 1, Name: "Makanan", Description: "Kategori makanan ringan"},
			{ID: 2, Name: "Minuman", Description: "Kategori minuman"},
		},
	}
}

func (r *categoryRepository) FindAll() []model.Category {
	return r.category
}

func (r *categoryRepository) FindByID(id int) (*model.Category, bool) {
	for _, c := range r.category {
		if c.ID == id {
			return &c, true
		}
	}
	return nil, false
}

func (r *categoryRepository) Create(category model.Category) model.Category {
	category.ID = len(r.category) + 1
	r.category = append(r.category, category)
	return category
}

func (r *categoryRepository) Update(id int, category model.Category) (*model.Category, bool) {
	for i := range r.category {
		if r.category[i].ID == id {
			category.ID = id
			r.category[i] = category
			return &category, true
		}
	}
	return nil, false
}

func (r *categoryRepository) Delete(id int) bool {
	for i, c := range r.category {
		if c.ID == id {
			r.category = append(r.category[:i], r.category[i+1:]...)
			return true
		}
	}
	return false
}
