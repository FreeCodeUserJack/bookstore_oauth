package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Errorf("expiration time should be 24 hours but got: %d", expirationTime)
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)

	if at.UserId != 1 {
		t.Fatal("expected not nil token but got 0 value for UserId")
	}

	if at.IsExpired() {
		t.Errorf("expected to not be expired but is expired with time of: %d while current time is: %d", at.Expires, time.Now().UTC().Unix())
	}

	if at.AccessToken != "" {
		t.Errorf("new access token should not have access token but got: %s", at.AccessToken)
	}

	if at.UserId != 0 {
		t.Errorf("new access token should be 0 but got: %d", at.UserId)
	}
}

func TestIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()
	if at.IsExpired() {
		t.Error("access token expiring 3 hours from now should not be expired")
	}
}