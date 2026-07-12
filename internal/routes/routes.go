package routes

import (
	"net/http"

	"github.com/cryskram/expense-tracker-go/internal/handlers"
	"github.com/cryskram/expense-tracker-go/internal/response"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, categoryHandler *handlers.CategoryHandler) {
	router.GET("/health", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "server is live", "")
	})

	api := router.Group("/api")
	{
		api.GET("/categories", categoryHandler.GetAll)
		api.POST("/categories", categoryHandler.Create)
	}

}
