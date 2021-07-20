package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/crypto_utils"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/errors"
)


const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AccessTokenRequest) IsValid() *errors.RestError {
	a.Email = strings.TrimSpace(a.Email)
	a.Password = strings.TrimSpace(a.Password)

	if a.Email == "" || a.Password == "" {
		return errors.NewBadRequestError("invalid credentials", "bad request")
	}

	return nil
}

func (a *AccessToken) IsValid() *errors.RestError {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
	if a.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id", "bad request")
	}
	if a.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id", "bad request")
	}
	if a.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id", "bad request")
	}
	if a.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time", "bad request") 
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(time.Hour * expirationTime).Unix(),
	}
}

func (a *AccessToken) IsExpired() bool {
	// can be one liner but this is more readable
	now := time.Now().UTC()
	expirationTime := time.Unix(a.Expires, 0)
	return expirationTime.Before(now)
}

func (a *AccessToken) Generate() {
	a.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", a.UserId, a.Expires))
}