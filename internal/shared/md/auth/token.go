package auth

import "time"

/* __________________________________________________ */

type TokenModel struct {
	Token     string
	ExpiresAt time.Time
}

type TokensModel struct {
	AccessToken  TokenModel
	RefreshToken TokenModel
}

/* __________________________________________________ */
