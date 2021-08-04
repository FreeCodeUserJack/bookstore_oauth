package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/crypto_utils"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
)


const (
	expirationTime = 24
	grantTypePassword = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	// used for password grant_type
	Username    string `json:"username"`
	Password string `json:"password"`

	// used for client_crdentials grant_type
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Scope string `json:"scope"`
}

func (a *AccessTokenRequest) IsValid() rest_errors.RestError {
	switch a.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant type parameter")
	}

	a.Username = strings.TrimSpace(a.Username)
	a.Password = strings.TrimSpace(a.Password)

	if a.Username == "" || a.Password == "" {
		return rest_errors.NewBadRequestError("invalid credentials")
	}

	return nil
}

func (a *AccessToken) IsValid() rest_errors.RestError {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
	if a.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if a.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if a.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if a.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time") 
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