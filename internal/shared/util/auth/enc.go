package auth

import (
	"time"

	timex "github.com/evgenivanovi/gpl/std/time"
	"github.com/golang-jwt/jwt/v5"
)

/* __________________________________________________ */

func CreateAccessToken(user User, exp time.Duration) *jwt.Token {

	return jwt.NewWithClaims(
		MethodProvider(),
		AccessClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					timex.NowPlus(exp),
				),
			},
			User: &user,
		},
	)

}

func CreateRefreshToken(user User, exp time.Duration) *jwt.Token {

	return jwt.NewWithClaims(
		MethodProvider(),
		RefreshClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					timex.NowPlus(exp),
				),
			},
			UserID: user.UserID,
		},
	)

}

/* __________________________________________________ */

func DecodeAccessClaims(token string, secret string) (*AccessClaims, error) {

	claims := new(AccessClaims)

	tkn, err := jwt.ParseWithClaims(
		token,
		claims,
		KeyProvider(secret),
	)

	if err != nil && !tkn.Valid {
		return nil, err
	}

	return claims, nil

}

func DecodeUser(token string, secret string) (*User, error) {

	claims, err := DecodeAccessClaims(token, secret)

	if err != nil {
		return nil, err
	}

	return claims.User, nil

}

/* __________________________________________________ */
