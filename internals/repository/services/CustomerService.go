package services

import (
	"fmt"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	"github.com/google/uuid"
)

type CustomerServices struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (cs *CustomerServices) AllCustomersService(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Customer, error) {
	customers, err := cs.PostgresDBRepo.AllCustomers(limit, offset, optionalParams)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (cs *CustomerServices) OneCustomerServiceByID(id string) (*models.Customer, error) {

	err := uuid.Validate(id)
	if err != nil {
		fmt.Printf("Error validating UUID: %v\n", err)
		return nil, err
	}

	customer, err := cs.PostgresDBRepo.OneCustomerByID(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
