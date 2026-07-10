package main

import (
	"log"

	"github.com/cryskram/expense-tracker-go/config"
	"github.com/cryskram/expense-tracker-go/internal/database"
	"github.com/cryskram/expense-tracker-go/internal/routes"
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
	routes.Register(router)

	log.Printf("Starting %s on port %s", cfg.APP_NAME, cfg.PORT)

	if err := router.Run(":" + cfg.PORT); err != nil {
		log.Fatal(err)
	}
}
