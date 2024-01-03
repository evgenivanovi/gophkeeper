package token

import (
	"time"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
)

/* __________________________________________________ */

type Token struct {
	Token     string
	ExpiresAt time.Time
}

func NewToken(token string, expiration time.Time) *Token {
	return &Token{
		Token:     token,
		ExpiresAt: expiration,
	}
}

/* __________________________________________________ */

type Tokens struct {
	AccessToken  Token
	RefreshToken Token
}

func NewTokens(access Token, refresh Token) *Tokens {
	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type TokenData struct {
	UserID common.UserID
}

func NewTokenData(userID common.UserID) *TokenData {
	return &TokenData{
		UserID: userID,
	}
}

/* __________________________________________________ */
