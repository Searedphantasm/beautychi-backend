package services

import (
	"errors"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserService struct {
	PostgresDBRepo *dbrepo.PostgresDBRepo
}

func (s *UserService) RegisterAdminUserService(userInfo models.User) error {

	err := validation.ValidateStruct(&userInfo,
		validation.Field(&userInfo.Username, validation.Required, validation.Length(1, 255)),
		validation.Field(&userInfo.Password, validation.Required, validation.Length(1, 255)),
	)

	if err != nil {
		return err
	}

	isExist, err := s.PostgresDBRepo.IsUserExist(userInfo.Username, userInfo.Phone)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("user with same credentials already exists")
	}

	// hashpassword
	encryptor := EncryptionService{Key: nil}
	hashedPassword, err := encryptor.NewHashedPassword(userInfo.Password)
	if err != nil {
		return err
	}

	userInfo.Password = hashedPassword
	err = s.PostgresDBRepo.RegisterAdminUser(userInfo)
	if err != nil {
		return err
	}

	return nil
}
