package services

import (
	"errors"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	apperrors "github.com/cryskram/expense-tracker-go/internal/errors"
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/cryskram/expense-tracker-go/internal/repositories"
	"github.com/cryskram/expense-tracker-go/internal/utils"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type CategoryService interface {
	Create(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetAll() ([]dto.CategoryResponse, error)
	GetByID(id uuid.UUID) (*dto.CategoryResponse, error)
	Update(id uuid.UUID, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error)
	Delete(id uuid.UUID) error
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

	resp := utils.ToCategoryResponse(category)
	return &resp, nil
}

func (s *categoryService) GetAll() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.CategoryResponse, 0, len(categories))

	for _, category := range categories {
		response = append(response, utils.ToCategoryResponse(category))
	}

	return response, nil
}

func (s *categoryService) GetByID(id uuid.UUID) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrCategoryNotFound
		}

		return nil, err
	}

	resp := utils.ToCategoryResponse(*category)
	return &resp, nil
}

func (s *categoryService) Update(id uuid.UUID, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Type = models.Type(req.Type)
	category.Icon = req.Icon
	category.Color = req.Color

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	resp := utils.ToCategoryResponse(*category)
	return &resp, nil
}

func (s *categoryService) Delete(id uuid.UUID) error {
	_, err := s.repo.GetByID(id)

	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
