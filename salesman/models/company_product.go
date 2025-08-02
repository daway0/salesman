package models

import (
	"github.com/google/uuid"
)

type CompanyProduct struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
}
