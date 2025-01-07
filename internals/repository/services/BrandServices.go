package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	validation "github.com/go-ozzo/ozzo-validation"
)

type BrandServices struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (bs *BrandServices) AllBrandsService() ([]*models.Brand, error) {
	brands, err := bs.PostgresDBRepo.AllBrands()
	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (bs *BrandServices) CreateBrandService(brand models.Brand) error {
	err := validation.ValidateStruct(&brand,
		validation.Field(&brand.Name, validation.Required),
		validation.Field(&brand.Description, validation.Required),
		validation.Field(&brand.Country, validation.Required),
		validation.Field(&brand.Logo, validation.Required),
		validation.Field(&brand.Slug, validation.Required))
	if err != nil {
		return err
	}

	err = bs.PostgresDBRepo.InsertBrand(brand)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BrandServices) DeleteBrandService(brandID int) error {

	err := bs.PostgresDBRepo.DeleteBrandByID(brandID)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BrandServices) UpdateBrandService(brand models.Brand) error {
	err := validation.ValidateStruct(&brand,
		validation.Field(&brand.Name, validation.Required),
		validation.Field(&brand.Description, validation.Required),
		validation.Field(&brand.Country, validation.Required),
		validation.Field(&brand.Slug, validation.Required),
		validation.Field(&brand.LogoKey, validation.Required),
		validation.Field(&brand.Logo, validation.Required),
		validation.Field(&brand.Website, validation.Required),
	)
	if err != nil {
		return err
	}

	err = bs.PostgresDBRepo.UpdateBrand(brand)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BrandServices) GetBrandServiceByIDOrSlug(brandID int, brandSlug string) (*models.Brand, error) {

	params := models.OneParams{
		ID:   brandID,
		Slug: brandSlug,
	}

	brand, err := bs.PostgresDBRepo.GetBrandByIDOrSlug(params)
	if err != nil {
		return nil, err
	}

	return brand, nil
}
