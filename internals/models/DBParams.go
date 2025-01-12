package models

type OneParams struct {
	ID   int
	Slug string
}

type OptionalQueryParams struct {
	Search          string
	Filter          string
	ProductCategory string
}
