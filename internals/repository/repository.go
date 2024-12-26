package repository

import (
	"database/sql"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	AllSubCategories() ([]*models.SubCategory, error)
	UpdateSubCategory(subCategory models.SubCategory) error
	GetSubCategoryByID(subCategoryID int) (*models.SubCategory, error)
	DeleteSubCategoryByID(subCategoryID int) error
	InsertSubCategory(subCategory models.SubCategory) error

	AllCategories() ([]*models.Category, error)
	InsertCategory(category models.Category) error
	UpdateCategoryByID(category models.Category) error
	DeleteCategoryByID(categoryID int) error
	GetCategoryByID(id int) (*models.Category, error)

	AllProducts(limit, offset int) ([]*models.Product, error)
	ProductByID(id int) (*models.Product, error)
	InsertProduct(product models.Product) error
	UpdateProduct(product models.Product) error
	UpdateProductImages(productID int, productImages []models.ProductImage) error

	AllBrands() ([]*models.Brand, error)
	InsertBrand(brand models.Brand) error
	DeleteBrandByID(brandID int) error
	GetBrandByID(brandID int) (*models.Brand, error)
	UpdateBrand(brand models.Brand) error
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

	GetCategoryByIDService(categoryID int) (*models.Category, error)
	DeleteCategoryService(categoryID int) error
}

type SubCategoryRepo interface {
	AllSubCategoryService() ([]*models.SubCategory, error)
	InsertSubCategoryService(subCategory models.SubCategory) error
	UpdateSubCategoryService(subCategory models.SubCategory) error
	DeleteSubCategoryService(subCategoryID int) error
}

type BrandRepo interface {
	AllBrandsService() ([]*models.Brand, error)
	DeleteBrandService(brandID int) error
	CreateBrandService(brand models.Brand) error
	UpdateBrandService(brand models.Brand) error
	GetBrandService(brandID int) (*models.Brand, error)
}
