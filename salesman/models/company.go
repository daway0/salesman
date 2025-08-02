package models

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID uuid.UUID `json:"id" db:"id"`

	Title string `json:"title" db:"title"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

