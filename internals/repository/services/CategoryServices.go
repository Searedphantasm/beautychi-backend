package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
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
