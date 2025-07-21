package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	CompanyID   uuid.UUID  `json:"company_id" db:"company_id"`
	Title       string     `json:"title" db:"title"`
	Description *string    `json:"description" db:"description"`
	Price       float64    `json:"price" db:"price"`
	ImageURL    *string    `json:"image_url" db:"image_url"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
