package auth

import "time"

/* __________________________________________________ */

type AccessTokenModel struct {
	Token     string
	ExpiresAt time.Time
}

type RefreshTokenModel struct {
	Token     string
	ExpiresAt time.Time
}

type TokensModel struct {
	AccessToken  AccessTokenModel
	RefreshToken RefreshTokenModel
}

/* __________________________________________________ */
