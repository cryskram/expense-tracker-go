package dto

type CreateTransactionRequest struct {
	Title           string  `json:"title" binding:"required,max=255"`
	Notes           string  `json:"notes"`
	Amount          float64 `json:"amount" binding:"required,gt=0"`
	CategoryID      string  `json:"category_id" binding:"required,uuid"`
	TransactionDate string  `json:"transaction_date" binding:"required"`
}

type TransactionResponse struct {
	ID              string           `json:"id"`
	Title           string           `json:"title"`
	Notes           string           `json:"notes"`
	Amount          float64          `json:"amount"`
	TransactionDate string           `json:"transaction_date"`
	Category        CategoryResponse `json:"category"`
}
