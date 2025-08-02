package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id" db:"id"`

	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	NSID        string     `json:"nsid" db:"nsid"`
	Birthdate   time.Time  `json:"birthdate" db:"birthdate"`
	Email       string     `json:"email" db:"email"`
	Password    string     `json:"password" db:"password"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UpdateUserRequest struct {
	FirstName *string    `json:"first_name"`
	LastName  *string    `json:"last_name"`
	NSID      *string    `json:"nsid"`
	Birthdate *time.Time `json:"birthdate"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
}

type CustomerCreate struct {
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	NSID        string     `json:"nsid" db:"nsid"`
}


type CompanyCustomerProductCreate struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	SubscriptionNo string `json:"subscription_no" db:"subscription_no"`
	Number string `json:"number" db:"number"`
}

type UserProductPath struct {
	ID uuid.UUID `json:"id" db:"id"`
	UserProductPath string `json:"user_product_path" db:"user_product_path"`
}

type ServicePath struct {
	ID uuid.UUID `json:"id" db:"id"`
	ServicePath string `json:"service_path" db:"service_path"`
}