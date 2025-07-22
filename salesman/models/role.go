package models

import (
	"github.com/google/uuid"
	"time"
)

type Role struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description *string    `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type UserRoleInfo struct {
	RoleID      uuid.UUID `json:"role_id"`
	Title       string    `json:"title"`
	Description *string    `json:"description"`
	UserHasRole bool      `json:"user_has_role"`
}
