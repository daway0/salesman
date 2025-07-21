package models

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Title       string         `json:"title" db:"title"`
	BrandName   *string `json:"brand_name" db:"brand_name"`
	CID         *string `json:"cid" db:"cid"`
	Description *string `json:"description" db:"description"`
	ImageURL    *string `json:"image_url" db:"image_url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time   `json:"deleted_at" db:"deleted_at"`
}
