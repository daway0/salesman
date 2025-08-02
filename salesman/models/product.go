package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID uuid.UUID `json:"id" db:"id"`

	CompanyID     uuid.UUID `json:"company_id" db:"company_id"`
	Title         string    `json:"title" db:"title"`
	HasCommission bool      `json:"has_commission" db:"has_commission"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type ProductWithCompany struct {
	Product
	CompanyTitle string `json:"company_title" db:"company_title"`
}
