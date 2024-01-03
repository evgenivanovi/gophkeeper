package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

/* __________________________________________________ */

type AccessClaims struct {
	jwt.RegisteredClaims
	User *User `json:"user"`
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

/* __________________________________________________ */
