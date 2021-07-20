package rest

import (
	"encoding/json"
	"time"

	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/users"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100*time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct {
}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (u *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestError) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user", "internal server error")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user", "internal server error")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error trying to umarshal users response", "internal server error")
	}

	return &user, nil
}