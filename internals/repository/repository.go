package repository

import (
	"bytes"
	"database/sql"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	AllSubCategories() ([]*models.SubCategory, error)
	UpdateSubCategory(subCategory models.SubCategory) error
	GetSubCategoryByIDOrSlug(params models.OneParams) (*models.SubCategory, error)
	DeleteSubCategoryByID(subCategoryID int) error
	InsertSubCategory(subCategory models.SubCategory) error

	AllCategories(optionalParams models.OptionalQueryParams) ([]*models.Category, error)
	InsertCategory(category models.Category) error
	UpdateCategoryByID(category models.Category) error
	DeleteCategoryByID(categoryID int) error
	GetCategoryByIDOrSlug(params models.OneParams) (*models.Category, error)

	AllProducts(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Product, error)
	ProductByIDOrSlug(params models.OneParams) (*models.Product, error)
	InsertProduct(product models.Product) error
	UpdateProduct(product models.Product) error
	UpdateProductImages(productID int, productImages []models.ProductImage) error

	AllBrands() ([]*models.Brand, error)
	InsertBrand(brand models.Brand) error
	DeleteBrandByID(brandID int) error
	GetBrandByIDOrSlug(params models.OneParams) (*models.Brand, error)
	UpdateBrand(brand models.Brand) error

	AllCustomers(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Customer, error)
	OneCustomerByID(id string) (*models.Customer, error)
}

type ProductRepo interface {
	AllProductsService(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Product, error)
	OneProductServiceByProductIDOrSlug(productID int, productSlug string) (*models.Product, error)

	InsertProductService(product models.Product) error
	UpdateProductService(product models.Product) error
	UpdateProductImagesService(productID int, productImages []models.ProductImage) error
	DeleteProductService(productID int) error
}

type CategoryRepo interface {
	AllCategoryService(optionalParams models.OptionalQueryParams) ([]*models.Category, error)
	CreateCategoryService(category models.Category) error
	UpdateCategoryService(category models.Category) error

	OneCategoryServiceByCategoryIDOrSlug(categoryID int, categorySlug string) (*models.Category, error)
	DeleteCategoryService(categoryID int) error
}

type SubCategoryRepo interface {
	AllSubCategoryService() ([]*models.SubCategory, error)
	OneSubCategoryServiceByIDOrSlug(subCategoryID int, subCategorySlug string) (*models.SubCategory, error)
	InsertSubCategoryService(subCategory models.SubCategory) error
	UpdateSubCategoryService(subCategory models.SubCategory) error
	DeleteSubCategoryService(subCategoryID int) error
}

type BrandRepo interface {
	AllBrandsService() ([]*models.Brand, error)
	DeleteBrandService(brandID int) error
	CreateBrandService(brand models.Brand) error
	UpdateBrandService(brand models.Brand) error
	GetBrandServiceByIDOrSlug(brandID int, brandSlug string) (*models.Brand, error)
}

type CustomerRepo interface {
	AllCustomersService(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Customer, error)
	OneCustomerServiceByID(id string) (*models.Customer, error)
}

type FileRepo interface {
	UploadUsingS3(fileContent *bytes.Reader, fileName string) (string, error)
}
