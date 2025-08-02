package models

import (
	"time"

	"github.com/google/uuid"
)

type CompanyUser struct {
	ID uuid.UUID `json:"id" db:"id"`

	CompanyID      uuid.UUID `json:"company_id" db:"company_id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	SubscriptionNo string    `json:"subscription_no" db:"subscription_no"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}


type CompanyUserProductList struct {
	CustomerCreate
	ID uuid.UUID `json:"id" db:"id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	ProductCompanyTitle string `json:"product_company_title" db:"product_company_title"`
	SubscriptionNo string `json:"subscription_no" db:"subscription_no"`
	Number string `json:"number" db:"number"`
	IsActive bool `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}