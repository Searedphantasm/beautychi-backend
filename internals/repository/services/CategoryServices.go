package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CategoryServices struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (cs *CategoryServices) AllCategoryService() ([]*models.Category, error) {
	categories, err := cs.PostgresDBRepo.AllCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryServices) CreateCategoryService(category models.Category) error {
	err := validation.ValidateStruct(&category,
		validation.Field(&category.Name, validation.Required),

		validation.Field(&category.Slug, validation.Required),

		validation.Field(&category.Description, validation.Required),

		validation.Field(&category.Image, validation.Required),

		validation.Field(&category.ImageKey, validation.Required),
	)

	if err != nil {
		return err
	}

	err = cs.PostgresDBRepo.InsertCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CategoryServices) UpdateCategoryService(category models.Category) error {
	err := validation.ValidateStruct(&category,
		validation.Field(&category.Name, validation.Required),
		validation.Field(&category.Slug, validation.Required),
		validation.Field(&category.Image, validation.Required),
		validation.Field(&category.ImageKey, validation.Required),
		validation.Field(&category.Description, validation.Required),
	)
	if err != nil {
		return err
	}

	err = cs.PostgresDBRepo.UpdateCategoryByID(category)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CategoryServices) DeleteCategoryService(categoryID int) error {
	err := cs.PostgresDBRepo.DeleteCategoryByID(categoryID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CategoryServices) OneCategoryServiceByCategoryIDOrSlug(categoryID int, categorySlug string) (*models.Category, error) {

	params := models.OneParams{
		ID:   categoryID,
		Slug: categorySlug,
	}

	category, err := cs.PostgresDBRepo.GetCategoryByIDOrSlug(params)
	if err != nil {
		return nil, err
	}

	return category, nil
}
