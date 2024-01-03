package auth

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	authmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	commonmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
)

/* __________________________________________________ */

type SigninUsecase interface {
	Execute(
		ctx context.Context, request SignInRequest,
	) (SignInResponse, error)
}

type SigninUsecaseService struct {
	trx     common.Transactor
	manager user.AuthManager
}

func ProvideSigninUsecaseService(
	trx common.Transactor,
	manager user.AuthManager,
) *SigninUsecaseService {
	return &SigninUsecaseService{
		trx:     trx,
		manager: manager,
	}
}

func (uc *SigninUsecaseService) Execute(
	ctx context.Context,
	request SignInRequest,
) (SignInResponse, error) {

	credentials := auth.NewCredentials(
		request.Payload.Credentials.Username,
		request.Payload.Credentials.Password,
	)

	ctx = uc.trx.StartEx(ctx)
	usr, err := uc.manager.Signin(ctx, credentials)
	uc.trx.CloseEx(ctx, err)

	if err != nil {
		return SignInResponse{}, err
	}

	return uc.toResponse(usr), nil

}

/* __________________________________________________ */

func (uc *SigninUsecaseService) toResponse(entity user.AuthUser) SignInResponse {
	return SignInResponse{
		Payload: SignInResponsePayload{
			Session: authmd.SessionModel{
				ID: entity.Data().SessionID.ID(),
				Tokens: authmd.TokensModel{
					AccessToken: authmd.AccessTokenModel{
						Token:     entity.Data().Tokens.AccessToken.Token,
						ExpiresAt: entity.Data().Tokens.AccessToken.ExpiresAt,
					},
					RefreshToken: authmd.RefreshTokenModel{
						Token:     entity.Data().Tokens.RefreshToken.Token,
						ExpiresAt: entity.Data().Tokens.RefreshToken.ExpiresAt,
					},
				},
			},
			User: authmd.UserModel{
				ID:       entity.Identity().ID(),
				Metadata: commonmd.FromMetadata(entity.Metadata()),
			},
		},
	}
}

/* __________________________________________________ */
