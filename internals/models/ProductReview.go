package models

import "github.com/jackc/pgx/v5/pgtype"

type ProductReview struct {
	ID               int         `json:"id"`
	ProductID        int         `json:"product_id"`
	CustomerID       pgtype.UUID `json:"customer_id,omitempty"`
	CustomerUsername string      `json:"customer_username,omitempty"`
	Rate             int         `json:"rate"`
	ReviewBody       string      `json:"review_body,omitempty"`
	Accepted         bool        `json:"accepted"`
}
