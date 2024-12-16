package repository

import (
	"database/sql"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllProducts() ([]*models.Product, error)
	AllCategories() ([]*models.Category, error)
	AllBrands() ([]*models.Brand, error)
}

type ProductRepo interface {
	AllProductsService() ([]*models.Product, error)
}

type CategoryRepo interface {
	AllCategoryService() ([]*models.Category, error)
}

type BrandRepo interface {
	AllBrandsService() ([]*models.Brand, error)
}
