package models

import "github.com/google/uuid"

type ContentType struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Model string    `json:"model" db:"model"`
}
