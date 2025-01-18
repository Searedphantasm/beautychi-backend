package dbrepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"log"
)

// TODO: registering admin user

func (r *PostgresDBRepo) RegisterAdminUser(userInfo models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string

	query = `INSERT INTO admin_user ( username, hashed_password, permission,phone) VALUES ( $1, $2, $3, $4);`
	log.Println(userInfo)
	_, err := r.DB.ExecContext(ctx, query, userInfo.Username, userInfo.HashedPassword, userInfo.Permission, userInfo.Phone)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) IsUserExist(username, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var exists bool
	err := r.DB.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM admin_user WHERE phone = $1 OR username = $2)", phone, username).Scan(&exists)
	if err != nil {
		return false, errors.New(fmt.Sprint("Error checking if user exists: ", err))
	}

	if exists {
		return true, nil
	}

	return false, nil
}
