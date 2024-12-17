package models

type ProductSpecification struct {
	ID               int    `json:"id"`
	ProductID        int    `json:"product_id"`
	SpecsTitle       string `json:"specs_title"`
	SpecsDescription string `json:"specs_description"`
}
