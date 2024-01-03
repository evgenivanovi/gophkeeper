package auth

import "time"

/* __________________________________________________ */

const TokenAccessSecretKey = "default"
const TokenAccessExpirationTime = time.Minute * 15

const TokenRefreshSecretKey = "default"
const TokenRefreshExpirationTime = time.Hour * 24 * 60

/* __________________________________________________ */

type TokenOp func(*TokenSettings)

func (op TokenOp) Join(other TokenOp) TokenOp {
	return func(opts *TokenSettings) {
		op(opts)
		other(opts)
	}
}

func (op TokenOp) And(ops ...TokenOp) TokenOp {
	return func(settings *TokenSettings) {
		op(settings)
		for _, op := range ops {
			op := op
			op(settings)
		}
	}
}

func WithAccessSecret(secret string) TokenOp {
	return func(settings *TokenSettings) {
		settings.accessSecret = secret
	}
}

func WithAccessSecretFn(fn func() string) TokenOp {
	return WithAccessSecret(fn())
}

func WithAccessExpiration(expiration time.Duration) TokenOp {
	return func(settings *TokenSettings) {
		settings.accessExpiration = expiration
	}
}

func WithAccessExpirationFn(fn func() time.Duration) TokenOp {
	return WithAccessExpiration(fn())
}

func WithRefreshSecret(secret string) TokenOp {
	return func(settings *TokenSettings) {
		settings.refreshSecret = secret
	}
}

func WithRefreshSecretFn(fn func() string) TokenOp {
	return WithRefreshSecret(fn())
}

func WithRefreshExpiration(expiration time.Duration) TokenOp {
	return func(settings *TokenSettings) {
		settings.refreshExpiration = expiration
	}
}

func WithRefreshExpirationFn(fn func() time.Duration) TokenOp {
	return WithRefreshExpiration(fn())
}

/* __________________________________________________ */

type TokenSettings struct {
	accessSecret      string
	refreshSecret     string
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func (t *TokenSettings) AccessSecret() string {
	return t.accessSecret
}

func (t *TokenSettings) RefreshSecret() string {
	return t.refreshSecret
}

func (t *TokenSettings) AccessExpiration() time.Duration {
	return t.accessExpiration
}

func (t *TokenSettings) RefreshExpiration() time.Duration {
	return t.refreshExpiration
}

func NewTokenSettings(ops ...TokenOp) *TokenSettings {
	settings := tokenSettings()
	for _, op := range ops {
		op(settings)
	}
	return settings
}

func tokenSettings() *TokenSettings {
	return &TokenSettings{
		accessSecret:      TokenAccessSecretKey,
		refreshSecret:     TokenRefreshSecretKey,
		accessExpiration:  TokenAccessExpirationTime,
		refreshExpiration: TokenRefreshExpirationTime,
	}
}

/* __________________________________________________ */
