package utils

import (
	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
)

func ToCategoryResponse(category models.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:    category.ID.String(),
		Name:  category.Name,
		Type:  string(category.Type),
		Icon:  category.Icon,
		Color: category.Color,
	}
}

func ToCategoryResponses(categories []models.Category) []dto.CategoryResponse {
	response := make([]dto.CategoryResponse, 0, len(categories))

	for _, category := range categories {
		response = append(response, ToCategoryResponse(category))
	}

	return response
}
