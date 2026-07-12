package models

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	Income  Type = "income"
	Expense Type = "expense"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Type      Type      `gorm:"size:20;not null"`
	Icon      string    `gorm:"size:100;not null"`
	Color     string    `gorm:"size:20;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Category) TableName() string {
	return "categories"
}
