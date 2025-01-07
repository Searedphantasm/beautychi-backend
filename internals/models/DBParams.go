package models

type OneParams struct {
	ID   int
	Slug string
}

type AllQueryParams struct {
	Search string
	Filter string
}
