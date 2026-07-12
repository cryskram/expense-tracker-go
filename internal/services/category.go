package services

import "github.com/cryskram/expense-tracker-go/internal/dto"

type CategoryService interface {
	Create(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
	GetAll() ([]dto.CreateCategoryResponse, error)
}
