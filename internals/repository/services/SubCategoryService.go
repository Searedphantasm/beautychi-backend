package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
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
