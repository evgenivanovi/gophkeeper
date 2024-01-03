package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

/* __________________________________________________ */

func KeyProvider(secret string) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}
}

func MethodProvider() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}

func ClaimsProvider() func() jwt.Claims {
	return func() jwt.Claims {
		return &AccessClaims{}
	}
}

/* __________________________________________________ */
