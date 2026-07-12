package handlers

import (
	"net/http"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/response"
	"github.com/cryskram/expense-tracker-go/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")

		return
	}

	category, err := h.service.Create(req)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")

		return
	}

	response.Success(c, http.StatusCreated, "category created", category)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")

		return
	}

	response.Success(c, http.StatusOK, "categories retrieved", categories)
}
