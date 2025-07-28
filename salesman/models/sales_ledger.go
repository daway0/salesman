package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type SalesLedger struct {
	ID              uuid.UUID        `json:"id" db:"id"`
	CustomerID      uuid.UUID        `json:"customer_id" db:"customer_id"`
	ServiceID       uuid.UUID        `json:"service_id" db:"service_id"`
	CreatedBy       uuid.UUID        `json:"created_by" db:"created_by"`
	RefererID       *uuid.UUID       `json:"referer_id" db:"referer_id"`
	Price           float64          `json:"price" db:"price"`
	Discount        *float64         `json:"discount" db:"discount"`
	PaymentMethod   string           `json:"payment_method" db:"payment_method"`
	TRN             string           `json:"trn" db:"trn"`
	WorkflowHistory *json.RawMessage `json:"workflow_history" db:"workflow_history"`
	ApprovedBy      *uuid.UUID       `json:"approved_by" db:"approved_by"`
	ApprovedAt      *time.Time       `json:"approved_at" db:"approved_at"`
	CancelledBy     *uuid.UUID       `json:"cancelled_by" db:"cancelled_by"`
	CancelledAt     *time.Time       `json:"cancelled_at" db:"cancelled_at"`
	Status          string           `json:"status" db:"status"`
	CreatedAt       time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time       `json:"deleted_at" db:"deleted_at"`

	ServiceTitle      string  `json:"service_title"`
	CompanyTitle      string  `json:"company_title"`
	CustomerFirstName string  `json:"customer_first_name"`
	CustomerLastName  string  `json:"customer_last_name"`
	CreatedByName     string  `json:"created_by_name"`
	CreatedByLastName string  `json:"created_by_last_name"`
	RefererName       *string `json:"referer_name"`
	RefererLastName   *string `json:"referer_last_name"`
}

type Workflow struct {
	Step      string    `json:"step"`
	ActionAt  time.Time `json:"action_at"`
	Comment   string    `json:"comment"`
	ActorId   uuid.UUID `json:"actor_id"`
	ActorName string    `json:"actor_name"`
}

type SalesLedgerApproval struct {
	ApproverID uuid.UUID `json:"approved_by"`
	ApprovedAt time.Time `json:"approved_at"`
	Comment    string    `json:"comment"`
}

type SalesLedgerRejection struct {
	RejectedBy uuid.UUID `json:"rejected_by"`
	RejectedAt time.Time `json:"rejected_at"`
	Reason     string    `json:"reason"`
}

type SalesLedgerCancellation struct {
	CancelledBy uuid.UUID `json:"cancelled_by"`
	CancelledAt time.Time `json:"cancelled_at"`
	Reason      string    `json:"reason"`
}

type SalesLedgerResend struct {
	Price         float64  `json:"price" db:"price"`
	Discount      *float64 `json:"discount" db:"discount"`
	PaymentMethod string   `json:"payment_method" db:"payment_method"`
	Comment       string   `json:"comment" db:"comment"`
}
