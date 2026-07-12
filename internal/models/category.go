package models

import (
	"time"

	"github.com/google/uuid"
)

type CategoryType string

const (
	IncomeCategory  CategoryType = "income"
	ExpenseCategory CategoryType = "expense"
)

type Category struct {
	ID        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string       `gorm:"size:100;not null"`
	Type      CategoryType `gorm:"size:20;not null"`
	Icon      string       `gorm:"size:100;not null"`
	Color     string       `gorm:"size:20;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Category) TableName() string {
	return "categories"
}
