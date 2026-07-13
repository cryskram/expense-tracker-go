package services

import (
	"math"
	"time"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
	"github.com/cryskram/expense-tracker-go/internal/repositories"
	"github.com/cryskram/expense-tracker-go/internal/utils"
	"github.com/google/uuid"
)

type TransactionService interface {
	Create(req dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
	GetAll(filter dto.TransactionFilter) ([]dto.TransactionResponse, dto.Pagination, error)
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

	resp := utils.ToTransactionResponse(*savedTransaction)
	return &resp, nil
}

func (s *transactionService) GetAll(filter dto.TransactionFilter) ([]dto.TransactionResponse, dto.Pagination, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 20
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	if filter.Sort == "" {
		filter.Sort = "date"
	}

	if filter.Order == "" {
		filter.Order = "desc"
	}

	transactions, total, err := s.transactionRepo.GetAll(filter)

	if err != nil {
		return nil, dto.Pagination{}, err
	}

	totalPages := int(math.Ceil(
		float64(total) /
			float64(filter.Limit),
	))

	pagination := dto.Pagination{
		Page:        filter.Page,
		Limit:       filter.Limit,
		Total:       total,
		TotalPages:  totalPages,
		HasNext:     filter.Page < totalPages,
		HasPrevious: filter.Page > 1,
	}
	return utils.ToTransactionResponses(transactions), pagination, nil
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

	resp := utils.ToTransactionResponse(*transaction)
	return &resp, nil
}
