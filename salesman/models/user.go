package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	FirstName   string       `json:"first_name" db:"first_name"`
	LastName    string       `json:"last_name" db:"last_name"`
	NSID        string       `json:"nsid" db:"nsid"`
	Birthdate   time.Time    `json:"birthdate" db:"birthdate"`
	Email       string       `json:"email" db:"email"`
	Password    string       `json:"password" db:"password"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	LastLoginAt *time.Time   `json:"last_login_at,omitempty" db:"last_login_at"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt   *time.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UpdateUserRequest struct {
	FirstName *string    `json:"first_name"`
	LastName  *string    `json:"last_name"`
	NSID      *string    `json:"nsid"`
	Birthdate *time.Time `json:"birthdate"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
}

