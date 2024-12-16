package services

import (
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
)

type ProductServices struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (ps *ProductServices) AllProductsService() ([]*models.Product, error) {
	products, err := ps.PostgresDBRepo.AllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
