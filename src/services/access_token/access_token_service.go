package access_token

import (
	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/access_token"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/repository/db"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/repository/rest"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type service struct {
	restUsersRepo rest.RestUserRepository
	dbRepo db.DbRepository
}

func NewService(usersRepo rest.RestUserRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo: dbRepo,
	}
}

func (s *service) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id", "bad request")
	}
	return s.dbRepo.GetById(id)
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestError) {
	if err := request.IsValid(); err != nil {
		return nil, err
	}

	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()
	
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := at.IsValid(); err != nil {
		return err
	}

	return s.dbRepo.UpdateExpirationTime(at)
}