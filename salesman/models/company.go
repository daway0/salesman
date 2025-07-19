package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Title       string         `json:"title" db:"title"`
	BrandName   sql.NullString `json:"brand_name" db:"brand_name"`
	CID         sql.NullString `json:"cid" db:"cid"`
	Description sql.NullString `json:"description" db:"description"`
	ImageURL    sql.NullString `json:"image_url" db:"image_url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at" db:"deleted_at"`
}
