package repository

import (
	"database/sql"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllBrands() ([]*models.Brand, error)
	AllSubCategories() ([]*models.SubCategory, error)

	AllCategories() ([]*models.Category, error)
	InsertCategory(category models.Category) error
	UpdateCategoryByID(category models.Category) error

	AllProducts(limit, offset int) ([]*models.Product, error)
	ProductByID(id int) (*models.Product, error)
	InsertProduct(product models.Product) error
	UpdateProduct(product models.Product) error
	UpdateProductImages(productID int, productImages []models.ProductImage) error
}

type ProductRepo interface {
	AllProductsService(limit, offset int) ([]*models.Product, error)
	OneProductServiceByProductID(productID int) (*models.Product, error)

	InsertProductService(product models.Product) error
	UpdateProductService(product models.Product) error
	UpdateProductImagesService(productID int, productImages []models.ProductImage) error
	DeleteProductService(productID int) error
}

type CategoryRepo interface {
	AllCategoryService() ([]*models.Category, error)
	CreateCategoryService(category models.Category) error
	UpdateCategoryService(category models.Category) error
}

type SubCategoryRepo interface {
	AllSubCategoryService() ([]*models.SubCategory, error)
}

type BrandRepo interface {
	AllBrandsService() ([]*models.Brand, error)
}
