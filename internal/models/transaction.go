package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title           string    `gorm:"size:255;not null"`
	Notes           string    `gorm:"type:text"`
	Amount          float64   `gorm:"type:numeric(12,2);not null"`
	CategoryID      uuid.UUID `gorm:"type:uuid;not null"`
	Category        Category  `gorm:"foreignKey:CategoryID"`
	TransactionDate string    `gorm:"type:date;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Transaction) TableName() string {
	return "transactions"
}
