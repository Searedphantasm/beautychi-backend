package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Customer struct {
	ID        pgtype.UUID        `json:"id"`
	Username  string             `json:"username"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Addresses []*CustomerAddress `json:"addresses"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CustomerAddress struct {
	ID         int         `json:"id"`
	CustomerID pgtype.UUID `json:"customer_id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Email      string      `json:"email"`
	Address    string      `json:"address"`
	PostalCode string      `json:"postal_code"`
	State      string      `json:"state"`
	City       string      `json:"city"`
	Phone      string      `json:"phone"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
