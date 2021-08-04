package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/users"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100*time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestError)
}

type usersRepository struct {
}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (u *usersRepository) LoginUser(email, password string) (*users.User, rest_errors.RestError) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}
	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error trying to umarshal users response", errors.New("json parsing error"))
	}

	return &user, nil
}