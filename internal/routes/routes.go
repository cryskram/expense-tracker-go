package routes

import (
	"net/http"

	"github.com/cryskram/expense-tracker-go/internal/handlers"
	"github.com/cryskram/expense-tracker-go/internal/response"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, categoryHandler *handlers.CategoryHandler, transactionHandler *handlers.TransactionHandler) {
	router.GET("/health", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "server is live", "")
	})

	api := router.Group("/api")
	{
		categories := api.Group("/categories")
		{
			categories.POST("", categoryHandler.Create)
			categories.GET("", categoryHandler.GetAll)
			categories.GET("/:id", categoryHandler.GetByID)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
		}

		transactions := api.Group("/transactions")
		{
			transactions.POST("", transactionHandler.Create)
			transactions.GET("", transactionHandler.GetAll)
			transactions.GET("/:id", transactionHandler.GetByID)
		}
	}

}
