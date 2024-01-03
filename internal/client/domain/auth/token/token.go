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
	) (token.Token, error)
	ProvideRefresh(
		ctx context.Context, user string,
	) (token.Token, error)
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
) (token.Token, error) {

	cfg, err := t.reader.Read(common.MustOptionsFromCtx(ctx).Config)
	if err != nil {
		return token.Token{}, err
	}

	op := config.NewConfigObjectOperations(cfg)

	sec, ok := op.GetSecrets(user)
	if !ok {
		return token.Token{}, errx.NewErrorWithEntityCode(
			token.ErrorTokenEntity, core.ErrorNotFoundCode,
		)
	}

	access := token.NewToken(sec.Access.Data, sec.Access.Expiration.Time)
	return *access, nil

}

func (t *TokenProviderService) ProvideRefresh(
	ctx context.Context,
	user string,
) (token.Token, error) {

	cfg, err := t.reader.Read(common.MustOptionsFromCtx(ctx).Config)
	if err != nil {
		return token.Token{}, err
	}

	op := config.NewConfigObjectOperations(cfg)

	sec, ok := op.GetSecrets(user)
	if !ok {
		return token.Token{}, errx.NewErrorWithEntityCode(
			token.ErrorTokenEntity, core.ErrorNotFoundCode,
		)
	}

	access := token.NewToken(sec.Access.Data, sec.Access.Expiration.Time)
	return *access, nil

}

/* __________________________________________________ */
