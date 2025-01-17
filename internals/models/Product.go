package models

import "time"

type Product struct {
	ID                   int                     `json:"id"`
	Title                string                  `json:"title"`
	Slug                 string                  `json:"slug"`
	Description          string                  `json:"description"`
	Poster               string                  `json:"poster"`
	Price                int                     `json:"price"`
	Rate                 int                     `json:"rate"`
	PosterKey            string                  `json:"poster_key"`
	CategoryID           int                     `json:"category_id"`
	CategoryName         string                  `json:"category_name"`
	CategorySlug         string                  `json:"category_slug"`
	BrandID              int                     `json:"brand_id"`
	BrandName            string                  `json:"brand_name"`
	ProductStock         int                     `json:"product_stock"`
	ProductDiscountPrice int                     `json:"product_discount_price"`
	SubCategoryID        int                     `json:"sub_category_id"`
	SubCategoryName      string                  `json:"sub_category_name"`
	ConsumerGuide        string                  `json:"consumer_guide"`
	Contact              string                  `json:"contact"`
	ProductSpecs         []*ProductSpecification `json:"product_specs,omitempty"`
	ProductImages        []*ProductImage         `json:"product_images,omitempty"`
	Status               string                  `json:"status"`
	CreatedAt            time.Time               `json:"created_at"`
	UpdatedAt            time.Time               `json:"updated_at"`
}
