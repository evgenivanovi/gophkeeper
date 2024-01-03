package token

import (
	"time"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
)

/* __________________________________________________ */

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
}

func NewAccessToken(token string, expiration time.Time) *AccessToken {
	return &AccessToken{
		Token:     token,
		ExpiresAt: expiration,
	}
}

type RefreshToken struct {
	Token     string
	ExpiresAt time.Time
}

func NewRefreshToken(token string, expiration time.Time) *RefreshToken {
	return &RefreshToken{
		Token:     token,
		ExpiresAt: expiration,
	}
}

/* __________________________________________________ */

type Tokens struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
}

func NewTokens(access AccessToken, refresh RefreshToken) *Tokens {
	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

/* __________________________________________________ */

type Data struct {
	UserID common.UserID
}

func NewTokenData(userID common.UserID) *Data {
	return &Data{
		UserID: userID,
	}
}

/* __________________________________________________ */
