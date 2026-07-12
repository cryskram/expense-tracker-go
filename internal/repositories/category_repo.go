package repositories

import (
	"github.com/google/uuid"

	"github.com/cryskram/expense-tracker-go/internal/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id uuid.UUID) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uuid.UUID) error
}
