package models

import "time"

type Brand struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
	Logo        string    `json:"logo"`
	LogoKey     string    `json:"logo_key"`
	Website     string    `json:"website_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
