package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(time.Hour * expirationTime).Unix(),
	}
}

func (a AccessToken) IsExpired() bool {
	// can be one liner but this is more readable
	now := time.Now().UTC()
	expirationTime := time.Unix(a.Expires, 0)
	return expirationTime.Before(now)
}