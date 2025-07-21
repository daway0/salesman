package models

import (
	"time"

	"github.com/google/uuid"
)

type LogAction string

const (
	LogActionRead    LogAction = "READ"
	LogActionWrite   LogAction = "WRITE"
	LogActionUpdate  LogAction = "UPDATE"
	LogActionDelete  LogAction = "DELETE"
	LogActionExecute LogAction = "EXECUTE"
)

type Permission struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	ContentType uuid.UUID    `json:"content_type" db:"content_type"`
	ActionType  LogAction    `json:"action_type" db:"action_type"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
