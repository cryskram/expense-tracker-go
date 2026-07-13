package repositories

import (
	"strings"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetAll(filter dto.TransactionFilter) ([]models.Transaction, int64, error)
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

func (r *transactionRepository) GetAll(filter dto.TransactionFilter) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	query := r.db.
		Model(&models.Transaction{}).
		Preload("Category")

	if filter.CategoryID != "" {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.StartDate != "" {
		query = query.Where("transaction_date >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		query = query.Where("transaction_date <= ?", filter.EndDate)
	}

	allowedSort := map[string]string{
		"date":   "transaction_date",
		"amount": "amount",
		"title":  "title",
	}

	column, ok := allowedSort[filter.Sort]
	if !ok {
		column = "transaction_date"
	}

	order := "DESC"

	if strings.ToLower(filter.Order) == "asc" {
		order = "ASC"
	}

	query = query.Order(column + " " + order)

	var total int64

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (filter.Page - 1) * filter.Limit

	query = query.
		Limit(filter.Limit).
		Offset(offset)

	err := query.Find(&transactions).Error

	return transactions, total, err
}

func (r *transactionRepository) GetByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Category").First(&transaction, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
