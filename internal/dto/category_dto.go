package dto

type CategoryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type CategoryUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
