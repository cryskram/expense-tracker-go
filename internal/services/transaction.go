package services

import (
	"time"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/cryskram/expense-tracker-go/internal/repositories"
	"github.com/google/uuid"
)

type TransactionService interface {
	Create(req dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
	GetAll() ([]dto.TransactionResponse, error)
	GetByID(id string) (*dto.TransactionResponse, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	categoryRepo    repositories.CategoryRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, categoryRepo repositories.CategoryRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		categoryRepo:    categoryRepo,
	}
}

func (s *transactionService) Create(req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil, err
	}

	_, err = s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, err
	}

	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)

	if err != nil {
		return nil, err
	}

	transaction := models.Transaction{
		Title:           req.Title,
		Notes:           req.Notes,
		Amount:          req.Amount,
		CategoryID:      categoryID,
		TransactionDate: transactionDate,
	}

	if err := s.transactionRepo.Create(&transaction); err != nil {
		return nil, err
	}

	savedTransaction, err := s.transactionRepo.GetByID(transaction.ID)
	if err != nil {
		return nil, err
	}

	response := dto.TransactionResponse{
		ID:     savedTransaction.ID.String(),
		Title:  savedTransaction.Title,
		Notes:  savedTransaction.Notes,
		Amount: savedTransaction.Amount,
		Category: dto.CategoryResponse{
			ID:    savedTransaction.Category.ID.String(),
			Name:  savedTransaction.Category.Name,
			Type:  string(savedTransaction.Category.Type),
			Icon:  savedTransaction.Category.Icon,
			Color: savedTransaction.Category.Color,
		},
		TransactionDate: savedTransaction.TransactionDate.Format("2006-01-02"),
	}

	return &response, nil
}

func (s *transactionService) GetAll() ([]dto.TransactionResponse, error) {
	transactions, err := s.transactionRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.TransactionResponse, 0, len(transactions))

	for _, transaction := range transactions {
		response = append(response, dto.TransactionResponse{
			ID:     transaction.ID.String(),
			Title:  transaction.Title,
			Notes:  transaction.Notes,
			Amount: transaction.Amount,
			Category: dto.CategoryResponse{
				ID:    transaction.Category.ID.String(),
				Name:  transaction.Category.Name,
				Type:  string(transaction.Category.Type),
				Icon:  transaction.Category.Icon,
				Color: transaction.Category.Color,
			},
			TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
		})
	}

	return response, nil
}

func (s *transactionService) GetByID(id string) (*dto.TransactionResponse, error) {
	transactionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	transaction, err := s.transactionRepo.GetByID(transactionID)
	if err != nil {
		return nil, err
	}

	response := dto.TransactionResponse{
		ID:     transaction.ID.String(),
		Title:  transaction.Title,
		Notes:  transaction.Notes,
		Amount: transaction.Amount,
		Category: dto.CategoryResponse{
			ID:    transaction.Category.ID.String(),
			Name:  transaction.Category.Name,
			Type:  string(transaction.Category.Type),
			Icon:  transaction.Category.Icon,
			Color: transaction.Category.Color,
		},
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
	}

	return &response, nil
}
