package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	CompanyID   uuid.UUID      `json:"company_id" db:"company_id"`
	Title       string         `json:"title" db:"title"`
	Description sql.NullString `json:"description" db:"description"`
	Price       float64        `json:"price" db:"price"`
	ImageURL    sql.NullString `json:"image_url" db:"image_url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at" db:"deleted_at"`
}
