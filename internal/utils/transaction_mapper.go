package utils

import (
	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/models"
)

func ToTransactionResponse(transaction models.Transaction) dto.TransactionResponse {
	return dto.TransactionResponse{
		ID:              transaction.ID.String(),
		Title:           transaction.Title,
		Notes:           transaction.Notes,
		Amount:          transaction.Amount,
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
		Category: dto.CategoryResponse{
			ID:    transaction.Category.ID.String(),
			Name:  transaction.Category.Name,
			Type:  string(transaction.Category.Type),
			Icon:  transaction.Category.Icon,
			Color: transaction.Category.Color,
		},
	}
}

func ToTransactionResponses(transactions []models.Transaction) []dto.TransactionResponse {
	response := make([]dto.TransactionResponse, 0, len(transactions))

	for _, transaction := range transactions {
		response = append(response, ToTransactionResponse(transaction))
	}

	return response
}
