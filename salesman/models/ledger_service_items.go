package models

import (
	"github.com/google/uuid"
	"time"
)

type LedgerServiceItems struct {
	ID uuid.UUID `json:"id" db:"id"`

	LedgerID  uuid.UUID `json:"ledger_id" db:"ledger_id"`
	ServiceID uuid.UUID `json:"service_id" db:"service_id"`
	Price     int       `json:"price" db:"price"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
