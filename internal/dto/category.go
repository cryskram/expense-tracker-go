package dto

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required,max=100"`
	Type  string `json:"type" binding:"required,oneof=income expense"`
	Icon  string `json:"icon" binding:"required,max=100"`
	Color string `json:"color" binding:"required,max=20"`
}

type CreateCategoryResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}
