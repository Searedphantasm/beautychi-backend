package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SubCategoryServices struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (scs *SubCategoryServices) AllSubCategoryService() ([]*models.SubCategory, error) {
	subCategories, err := scs.PostgresDBRepo.AllSubCategories()
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}

func (scs *SubCategoryServices) InsertSubCategoryService(subCategory models.SubCategory) error {

	err := validation.ValidateStruct(&subCategory,
		validation.Field(&subCategory.Name, validation.Required),
		validation.Field(&subCategory.Description, validation.Required),
		validation.Field(&subCategory.ParentCategoryID, validation.Required),
		validation.Field(&subCategory.Image, validation.Required),
		validation.Field(&subCategory.Slug, validation.Required),
	)

	if err != nil {
		return err
	}

	err = scs.PostgresDBRepo.InsertSubCategory(subCategory)
	if err != nil {
		return err
	}

	return nil
}

func (scs *SubCategoryServices) UpdateSubCategoryService(subCategory models.SubCategory) error {
	err := validation.ValidateStruct(&subCategory,
		validation.Field(&subCategory.Name, validation.Required),
		validation.Field(&subCategory.Description, validation.Required),
		validation.Field(&subCategory.ParentCategoryID, validation.Required),
		validation.Field(&subCategory.Image, validation.Required),
		validation.Field(&subCategory.Slug, validation.Required),
	)

	if err != nil {
		return err
	}

	err = scs.PostgresDBRepo.UpdateSubCategory(subCategory)
	if err != nil {
		return err
	}

	return nil
}

func (scs *SubCategoryServices) DeleteSubCategoryService(subCategoryID int) error {
	err := scs.PostgresDBRepo.DeleteSubCategoryByID(subCategoryID)
	if err != nil {
		return nil
	}

	return nil
}

func (scs *SubCategoryServices) OneSubCategoryServiceByIDOrSlug(subCategoryID int, subCategorySlug string) (*models.SubCategory, error) {

	params := models.OneParams{
		ID:   subCategoryID,
		Slug: subCategorySlug,
	}

	subCategory, err := scs.PostgresDBRepo.GetSubCategoryByIDOrSlug(params)
	if err != nil {
		return nil, err
	}

	return subCategory, nil
}
