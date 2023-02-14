package repository

import "ZakirAvrora/go_test_backend/service1/internal/model"

type Repository interface {
	CreateUser(model.DbUser) error
	GetUser(email string) (*model.DbUser, error)
}
