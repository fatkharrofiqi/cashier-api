package repository

import (
	"cashier-api/internal/model"
	"database/sql"
	"fmt"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id int) (*model.Product, error)
	Create(product model.Product) (*model.Product, error)
	Update(id int, product model.Product) (*model.Product, error)
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.id, c.name, c.description
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		ORDER BY p.id
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		var categoryID sql.NullInt64
		var categoryIDFromJoin sql.NullInt64
		var categoryName sql.NullString
		var categoryDescription sql.NullString
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &categoryIDFromJoin, &categoryName, &categoryDescription); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		if categoryID.Valid && categoryIDFromJoin.Valid {
			p.CategoryID = new(int)
			*p.CategoryID = int(categoryID.Int64)
			p.Category = &model.Category{
				ID:          int(categoryIDFromJoin.Int64),
				Name:        categoryName.String,
				Description: categoryDescription.String,
			}
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

func (r *productRepository) FindByID(id int) (*model.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.id, c.name, c.description
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`
	var p model.Product
	var categoryID sql.NullInt64
	var category model.Category
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &category.ID, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query product: %w", err)
	}
	if categoryID.Valid {
		p.Category = &category
	}
	return &p, nil
}

func (r *productRepository) Create(product model.Product) (*model.Product, error) {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}
	return &product, nil
}

func (r *productRepository) Update(id int, product model.Product) (*model.Product, error) {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("product not found")
	}

	product.ID = id
	return &product, nil
}

func (r *productRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}
