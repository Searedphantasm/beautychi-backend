package models

import (
	"time"
)

type ProductImage struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Url       string    `json:"url"`
	UrlKey    string    `json:"url_key"`
	AltText   string    `json:"alt_text"`
	CreatedAt time.Time `json:"created_at"`
}
