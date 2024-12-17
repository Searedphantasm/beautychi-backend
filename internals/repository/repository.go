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
	AllSubCategories() ([]*models.SubCategory, error)

	ProductByID(id int) (*models.Product, error)
	InsertProduct(product models.Product) error
}

type ProductRepo interface {
	AllProductsService() ([]*models.Product, error)
	OneProductServiceByProductID(productID int) (*models.Product, error)

	InsertProductService(product models.Product) error
}

type CategoryRepo interface {
	AllCategoryService() ([]*models.Category, error)
}

type SubCategoryRepo interface {
	AllSubCategoryService() ([]*models.SubCategory, error)
}

type BrandRepo interface {
	AllBrandsService() ([]*models.Brand, error)
}
