package models

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID             pgtype.UUID
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	HashedPassword string `json:"-"`
	Permission     string `json:"permission"`
	CreateAt       string `json:"create_at"`
	UpdatedAt      string `json:"updated_at"`
}
