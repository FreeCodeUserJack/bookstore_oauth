package db

import (
	"github.com/FreeCodeUserJack/bookstore_oauth/src/clients/cassandra"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/access_token"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
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
	GetById(string) (*access_token.AccessToken, rest_errors.RestError)
	Create(access_token.AccessToken) rest_errors.RestError
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestError
}

type dbRepository struct {
}

func (d *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestError) {
	

	var res access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&res.AccessToken, &res.UserId, &res.ClientId, &res.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewBadRequestError(err.Error())
		}
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}

	return &res, nil
}

func (d *dbRepository) Create(ac access_token.AccessToken) rest_errors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, ac.AccessToken, ac.UserId, ac.ClientId, ac.Expires).Exec(); err != nil {
		return rest_errors.NewInternalServerError("internal server error", err)
	}

	return nil
}

func (d *dbRepository) UpdateExpirationTime(ac access_token.AccessToken) rest_errors.RestError {	
	if err := cassandra.GetSession().Query(queryUpdateExpires, ac.Expires, ac.AccessToken).Exec(); err != nil {
		return rest_errors.NewInternalServerError("internal server error", err)
	}

	return nil
}