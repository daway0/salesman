package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID               uuid.UUID       `json:"id" db:"id"`
	UserID           uuid.UUID       `json:"user_id" db:"user_id"`
	RoleID           uuid.UUID       `json:"role_id" db:"role_id"`
	PermissionID     uuid.UUID       `json:"permission_id" db:"permission_id"`
	RoleString       string          `json:"role_string" db:"role_string"`
	PermissionString string          `json:"permission_string" db:"permission_string"`
	Route            string          `json:"route" db:"route"`
	RequestPayload   json.RawMessage `json:"request_payload" db:"request_payload"`
	AdditionalData   json.RawMessage `json:"additional_data" db:"additional_data"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
}
