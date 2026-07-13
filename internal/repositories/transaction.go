package repositories

import (
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetAll() ([]models.Transaction, error)
	GetByID(id uuid.UUID) (*models.Transaction, error)
	// Update(transaction *models.Transaction) error
	// Delete(id uuid.UUID) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Category").Order("transaction_date DESC").Find(&transactions).Error

	return transactions, err
}

func (r *transactionRepository) GetByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Category").First(&transaction, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
