package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID uuid.UUID `json:"id" db:"id"`

	UserID uuid.UUID `json:"user_id" db:"user_id"`
	RoleID uuid.UUID `json:"role_id" db:"role_id"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
