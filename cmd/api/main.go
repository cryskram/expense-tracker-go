package main

import (
	"log"

	"github.com/cryskram/expense-tracker-go/config"
	"github.com/cryskram/expense-tracker-go/internal/database"
	"github.com/cryskram/expense-tracker-go/internal/handlers"
	"github.com/cryskram/expense-tracker-go/internal/repositories"
	"github.com/cryskram/expense-tracker-go/internal/routes"
	"github.com/cryskram/expense-tracker-go/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	_ = db

	router := gin.Default()
	categoryRepo := repositories.NewCategoryRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	categoryService := services.NewCategoryService(categoryRepo)
	transactionService := services.NewTransactionService(transactionRepo, categoryRepo)

	categoryHandler := handlers.NewCategoryHandler(categoryService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	routes.Register(router, categoryHandler, transactionHandler)

	log.Printf("Starting %s on port %s", cfg.APP_NAME, cfg.PORT)

	if err := router.Run(":" + cfg.PORT); err != nil {
		log.Fatal(err)
	}
}
