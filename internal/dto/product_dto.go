package dto

type ProductCreateRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID *int   `json:"category_id,omitempty"`
}

type ProductUpdateRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID *int   `json:"category_id,omitempty"`
}

type ProductResponse struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Price      int          `json:"price"`
	Stock      int          `json:"stock"`
	CategoryID *int         `json:"category_id"`
	Category   *CategoryDTO `json:"category,omitempty"`
}

type CategoryDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
