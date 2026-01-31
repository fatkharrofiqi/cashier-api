package repository

import (
	"cashier-api/internal/model"
	"database/sql"
	"fmt"
)

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
	FindByID(id int) (*model.Category, error)
	Create(category model.Category) (*model.Category, error)
	Update(id int, category model.Category) (*model.Category, error)
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll() ([]model.Category, error) {
	query := "SELECT id, name, description FROM categories ORDER BY id"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return categories, nil
}

func (r *categoryRepository) FindByID(id int) (*model.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	var c model.Category
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("category not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query category: %w", err)
	}
	return &c, nil
}

func (r *categoryRepository) Create(category model.Category) (*model.Category, error) {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}
	return &category, nil
}

func (r *categoryRepository) Update(id int, category model.Category) (*model.Category, error) {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := r.db.Exec(query, category.Name, category.Description, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("category not found")
	}

	category.ID = id
	return &category, nil
}

func (r *categoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}
