package services

import (
	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/cryskram/expense-tracker-go/internal/repositories"
	"github.com/google/uuid"
)

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error) {
	category := models.Category{
		ID:    uuid.New(),
		Name:  req.Name,
		Type:  models.CategoryType(req.Type),
		Icon:  req.Icon,
		Color: req.Color,
	}

	if err := s.repo.Create(&category); err != nil {
		return nil, err
	}

	return &dto.CreateCategoryResponse{
		ID:    category.ID.String(),
		Name:  category.Name,
		Type:  string(category.Type),
		Icon:  category.Icon,
		Color: category.Color,
	}, nil
}

func (s *categoryService) GetAll() ([]dto.CreateCategoryResponse, error) {
	categories, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	response := make([]dto.CreateCategoryResponse, 0, len(categories))

	for _, category := range categories {
		response = append(response, dto.CreateCategoryResponse{
			ID:    category.ID.String(),
			Name:  category.Name,
			Type:  string(category.Type),
			Icon:  category.Icon,
			Color: category.Color,
		})
	}

	return response, nil
}
