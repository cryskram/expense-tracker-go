package handlers

import (
	"net/http"

	"github.com/cryskram/expense-tracker-go/internal/dto"
	"github.com/cryskram/expense-tracker-go/internal/response"
	"github.com/cryskram/expense-tracker-go/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid category ID")
		return
	}

	category, err := h.service.GetByID(id)

	if err != nil {
		response.Error(c, http.StatusNotFound, "category not found")
		return
	}

	response.Success(c, http.StatusOK, "category retrieved", category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid category ID")
		return
	}

	var req dto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	category, err := h.service.Update(id, req)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}

	response.Success(c, http.StatusOK, "category updated", category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid category ID")
		return
	}

	err = h.service.Delete(id)

	if err != nil {
		response.Error(c, http.StatusNotFound, "category not found")
		return
	}

	c.Status(http.StatusNoContent)
}
