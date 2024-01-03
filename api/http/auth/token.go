package auth

import "time"

/* __________________________________________________ */

type AccessTokenModel struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type RefreshTokenModel struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type TokensModel struct {
	AccessToken  AccessTokenModel  `json:"accessToken"`
	RefreshToken RefreshTokenModel `json:"refreshToken"`
}

/* __________________________________________________ */
