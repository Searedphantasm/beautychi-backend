package models

import "time"

type Product struct {
	ID                   int64     `json:"id"`
	Title                string    `json:"title"`
	Slug                 string    `json:"slug"`
	Description          string    `json:"description"`
	Poster               string    `json:"poster"`
	Price                int64     `json:"price"`
	PosterKey            string    `json:"poster_key"`
	CategoryID           int64     `json:"category_id"`
	CategoryName         string    `json:"category_name"`
	BrandID              int64     `json:"brand_id"`
	BrandName            string    `json:"brand_name"`
	ProductStock         int64     `json:"product_stock"`
	ProductDiscountPrice int64     `json:"product_discount_price"`
	SubCategoryID        int64     `json:"sub_category_id"`
	SubCategoryName      string    `json:"sub_category_name"`
	ConsumerGuide        string    `json:"consumer_guide"`
	Contact              string    `json:"contact"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
