package token

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	errx "github.com/evgenivanovi/gpl/err"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type TokenProvider interface {
	ProvideAccess(
		ctx context.Context, user string,
	) (token.AccessToken, error)
	ProvideRefresh(
		ctx context.Context, user string,
	) (token.RefreshToken, error)
}

//goland:noinspection GoNameStartsWithPackageName
type TokenProviderService struct {
	reader config.Reader
}

func ProvideTokenProviderService(reader config.Reader) *TokenProviderService {
	return &TokenProviderService{
		reader: reader,
	}
}

func (t *TokenProviderService) ProvideAccess(
	ctx context.Context,
	user string,
) (token.AccessToken, error) {

	cfg, err := t.reader.Read(common.MustOptionsFromCtx(ctx).Config)
	if err != nil {
		return token.AccessToken{}, err
	}

	op := config.NewConfigObjectOperations(cfg)

	sec, ok := op.GetSecrets(user)
	if !ok {
		return token.AccessToken{}, errx.NewErrorWithEntityCode(
			token.ErrorTokenEntity, core.ErrorNotFoundCode,
		)
	}

	access := token.NewAccessToken(sec.Access.Data, sec.Access.Expiration.Time)
	return *access, nil

}

func (t *TokenProviderService) ProvideRefresh(
	ctx context.Context,
	user string,
) (token.RefreshToken, error) {

	cfg, err := t.reader.Read(common.MustOptionsFromCtx(ctx).Config)
	if err != nil {
		return token.RefreshToken{}, err
	}

	op := config.NewConfigObjectOperations(cfg)

	sec, ok := op.GetSecrets(user)
	if !ok {
		return token.RefreshToken{}, errx.NewErrorWithEntityCode(
			token.ErrorTokenEntity, core.ErrorNotFoundCode,
		)
	}

	access := token.NewRefreshToken(sec.Access.Data, sec.Access.Expiration.Time)
	return *access, nil

}

/* __________________________________________________ */
