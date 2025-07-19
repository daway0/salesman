package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

type SalesLedger struct {
	ID            uuid.UUID       `json:"id" db:"id"`
	ApprovedBy    NullUUID        `json:"approved_by" db:"approved_by"`
	ServiceID     uuid.UUID       `json:"service_id" db:"service_id"`
	Price         float64         `json:"price" db:"price"`
	SalesPrice    float64         `json:"sales_price" db:"sales_price"`
	SalesDiscount sql.NullFloat64 `json:"sales_discount" db:"sales_discount"`
	TRN           sql.NullString  `json:"trn" db:"trn"`
	ApprovedAt    sql.NullTime    `json:"approved_at" db:"approved_at"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     sql.NullTime    `json:"updated_at" db:"updated_at"`
	DeletedAt     sql.NullTime    `json:"deleted_at" db:"deleted_at"`
}
