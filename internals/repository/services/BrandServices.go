package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
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
