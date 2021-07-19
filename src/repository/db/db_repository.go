package db

import (
	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/access_token"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/errors"
)

func NewRepo() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func (d *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	// TODO: get access token from Cassandra ID
	return nil, errors.NewInternalServerError("db conn not impl", "internal server error")
}