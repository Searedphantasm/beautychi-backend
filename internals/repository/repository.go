package repository

import (
	"database/sql"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllProducts() ([]*models.Product, error)
}

type ProductRepo interface {
	AllProductsService() ([]*models.Product, error)
}
