package token

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	authtool "github.com/evgenivanovi/gophkeeper/internal/shared/util/auth"
	"github.com/evgenivanovi/gpl/std/time"
	"github.com/evgenivanovi/gpl/stdx/jwtx"
)

/* __________________________________________________ */

type Manager interface {
	Generate(ctx context.Context, data token.Data) token.Tokens
}

type ManagerService struct {
	settings authtool.TokenSettings
}

func ProvideManagerService(
	settings authtool.TokenSettings,
) *ManagerService {
	return &ManagerService{
		settings: settings,
	}
}

func (t *ManagerService) Generate(
	ctx context.Context, data token.Data,
) token.Tokens {

	accessToken := CreateAccessToken(data, t.settings.AccessExpiration())
	accessTokenString, _ := jwtx.SignJWT(*accessToken, t.settings.AccessSecret())
	access := token.NewAccessToken(accessTokenString, time.NowPlus(t.settings.AccessExpiration()))

	refreshToken := CreateRefreshToken(data, t.settings.RefreshExpiration())
	refreshTokenString, _ := jwtx.SignJWT(*refreshToken, t.settings.RefreshSecret())
	refresh := token.NewRefreshToken(refreshTokenString, time.NowPlus(t.settings.RefreshExpiration()))

	return *token.NewTokens(*access, *refresh)

}

/* __________________________________________________ */
