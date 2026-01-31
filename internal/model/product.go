package model

type Product struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	Price      int       `db:"price"`
	Stock      int       `db:"stock"`
	CategoryID *int      `db:"category_id"`
	Category   *Category
}

