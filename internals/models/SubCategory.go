package models

import "time"

type SubCategory struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	ParentCategoryID   int64     `json:"parent_category_id"`
	ParentCategoryName string    `json:"parent_category_name,omitempty"`
	Description        string    `json:"description"`
	Image              string    `json:"image"`
	ImageKey           string    `json:"image_key"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
