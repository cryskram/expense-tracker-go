package services

import (
	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/cryskram/expense-tracker-go/internal/repositories"
)

type CategoryService interface {
	Create(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetAll() ([]dto.CategoryResponse, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) Create(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:  req.Name,
		Type:  models.Type(req.Type),
		Icon:  req.Icon,
		Color: req.Color,
	}

	if err := s.repo.Create(&category); err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:    category.ID.String(),
		Name:  category.Name,
		Type:  string(category.Type),
		Icon:  category.Icon,
		Color: category.Color,
	}, nil
}

func (s *categoryService) GetAll() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.CategoryResponse, 0, len(categories))

	for _, category := range categories {
		response = append(response, dto.CategoryResponse{
			ID:    category.ID.String(),
			Name:  category.Name,
			Type:  string(category.Type),
			Icon:  category.Icon,
			Color: category.Color,
		})
	}

	return response, nil
}
