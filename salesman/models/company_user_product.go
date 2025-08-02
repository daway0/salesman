package models

import (
	"github.com/google/uuid"
	"time"
)

type CompanyUserProduct struct {
	ID uuid.UUID `json:"id" db:"id"`

	CompanyUserID uuid.UUID `json:"company_user_id" db:"company_user_id"`
	ProductID     string    `json:"product_id" db:"product_id"`
	IsActive      bool      `json:"is_active" db:"is_active"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

