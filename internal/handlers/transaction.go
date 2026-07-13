package handlers

import (
	"net/http"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/response"
	"github.com/cryskram/expense-tracker-go/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var req dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	transaction, err := h.service.Create(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}

	response.Success(c, http.StatusCreated, "transaction created", transaction)
}

func (h *TransactionHandler) GetAll(c *gin.Context) {

	var filter dto.TransactionFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid query parameters")
		return
	}

	transactions, pagination, err := h.service.GetAll(filter)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}

	response.Paginated(
		c,
		http.StatusOK,
		"transactions retrieved",
		transactions,
		pagination,
	)
}

func (h *TransactionHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid transaction ID")
		return
	}

	transaction, err := h.service.GetByID(id.String())
	if err != nil {
		response.Error(c, http.StatusNotFound, "transaction not found")
		return
	}

	response.Success(c, http.StatusOK, "transaction retrieved", transaction)
}
