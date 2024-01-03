package auth

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	authmdshared "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	commonmdshared "github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
)

/* __________________________________________________ */

type SignupUsecase interface {
	Execute(
		ctx context.Context, request SignUpRequest,
	) (SignUpResponse, error)
}

type SignupUsecaseService struct {
	trx     common.Transactor
	manager user.AuthManager
}

func ProvideSignupUsecaseService(
	trx common.Transactor,
	manager user.AuthManager,
) *SignupUsecaseService {
	return &SignupUsecaseService{
		trx:     trx,
		manager: manager,
	}
}

func (uc *SignupUsecaseService) Execute(
	ctx context.Context,
	request SignUpRequest,
) (SignUpResponse, error) {

	credentials := auth.NewCredentials(
		request.Payload.Credentials.Username,
		request.Payload.Credentials.Password,
	)

	ctx = uc.trx.StartEx(ctx)
	usr, err := uc.manager.SignupAndSignin(ctx, credentials)
	uc.trx.CloseEx(ctx, err)

	if err != nil {
		return SignUpResponse{}, err
	}

	return uc.toResponse(usr), nil

}

/* __________________________________________________ */

func (uc *SignupUsecaseService) toResponse(entity user.AuthUser) SignUpResponse {
	return SignUpResponse{
		Payload: SignUpResponsePayload{
			Session: authmdshared.SessionModel{
				ID: entity.Data().SessionID.ID(),
				Tokens: authmdshared.TokensModel{
					AccessToken: authmdshared.TokenModel{
						Token:     entity.Data().Tokens.AccessToken.Token,
						ExpiresAt: entity.Data().Tokens.AccessToken.ExpiresAt,
					},
					RefreshToken: authmdshared.TokenModel{
						Token:     entity.Data().Tokens.RefreshToken.Token,
						ExpiresAt: entity.Data().Tokens.RefreshToken.ExpiresAt,
					},
				},
			},
			User: authmdshared.UserModel{
				ID:       entity.Identity().ID(),
				Metadata: commonmdshared.FromMetadata(entity.Metadata()),
			},
		},
	}
}

/* __________________________________________________ */
