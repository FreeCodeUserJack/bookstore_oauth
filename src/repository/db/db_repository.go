package db

import (
	"github.com/FreeCodeUserJack/bookstore_oauth/src/clients/cassandra"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/access_token"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires = ? WHERE access_token=?;"
)

func NewRepo() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func (d *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	

	var res access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&res.AccessToken, &res.UserId, &res.ClientId, &res.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewBadRequestError(err.Error(), "access_token_id is not found")
		}
		return nil, errors.NewInternalServerError(err.Error(), "internal server error")
	}

	return &res, nil
}

func (d *dbRepository) Create(ac access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, ac.AccessToken, ac.UserId, ac.ClientId, ac.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error(), "internal server error")
	}

	return nil
}

func (d *dbRepository) UpdateExpirationTime(ac access_token.AccessToken) *errors.RestError {	
	if err := cassandra.GetSession().Query(queryUpdateExpires, ac.Expires, ac.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error(), "internal server error")
	}

	return nil
}