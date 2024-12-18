package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ProductServices struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (ps *ProductServices) AllProductsService(limit, offset int) ([]*models.Product, error) {
	products, err := ps.PostgresDBRepo.AllProducts(limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductServices) DeleteProductService(productID int) error {
	err := ps.PostgresDBRepo.DeleteProductByID(productID)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductServices) OneProductServiceByProductID(productID int) (*models.Product, error) {
	product, err := ps.PostgresDBRepo.ProductByID(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductServices) InsertProductService(product models.Product) error {

	err := validation.ValidateStruct(&product,

		validation.Field(&product.Title, validation.Required, validation.Length(1, 255)),

		validation.Field(&product.Description, validation.Required, validation.Length(1, 64000)),

		validation.Field(&product.Price, validation.Required, validation.Min(0)),

		validation.Field(&product.CategoryID, validation.Required, validation.Min(0)),

		validation.Field(&product.BrandID, validation.Required, validation.Min(0)),

		validation.Field(&product.SubCategoryID, validation.Required, validation.Min(0)),

		validation.Field(&product.ConsumerGuide, validation.Required, validation.Length(1, 64000)),

		validation.Field(&product.ProductStock, validation.Required, validation.Min(1)),

		validation.Field(&product.ProductDiscountPrice, validation.Required, validation.Min(0)),

		validation.Field(&product.Slug, validation.Required, validation.Length(1, 255)),

		validation.Field(&product.Contact, validation.Required),

		validation.Field(&product.Poster, validation.Required),

		validation.Field(&product.PosterKey, validation.Required),

		validation.Field(&product.Status, validation.Required, validation.In("Active", "Inactive")),

		validation.Field(&product.ProductSpecs, validation.Required),

		validation.Field(&product.ProductImages, validation.Required),
	)

	if err != nil {
		return err
	}

	err = ps.PostgresDBRepo.InsertProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductServices) UpdateProductService(product models.Product) error {
	err := validation.ValidateStruct(&product,
		validation.Field(&product.Title, validation.Required, validation.Length(1, 255)),

		validation.Field(&product.Description, validation.Required, validation.Length(1, 64000)),

		validation.Field(&product.Price, validation.Required, validation.Min(0)),

		validation.Field(&product.CategoryID, validation.Required, validation.Min(0)),

		validation.Field(&product.BrandID, validation.Required, validation.Min(0)),

		validation.Field(&product.SubCategoryID, validation.Required, validation.Min(0)),

		validation.Field(&product.ConsumerGuide, validation.Required, validation.Length(1, 64000)),

		validation.Field(&product.ProductStock, validation.Required, validation.Min(1)),

		validation.Field(&product.ProductDiscountPrice, validation.Required, validation.Min(0)),

		validation.Field(&product.Slug, validation.Required, validation.Length(1, 255)),

		validation.Field(&product.Contact, validation.Required),

		validation.Field(&product.Poster, validation.Required),

		validation.Field(&product.PosterKey, validation.Required),

		validation.Field(&product.Status, validation.Required, validation.In("Active", "Inactive")),
	)

	if err != nil {
		return err
	}

	err = ps.PostgresDBRepo.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
