package database

import (
	"log"

	"github.com/cryskram/expense-tracker-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.GetDBConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect db")
	}

	log.Println("connected to db")
	return db, nil
}
