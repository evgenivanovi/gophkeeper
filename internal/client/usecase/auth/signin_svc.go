package auth

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	configuc "github.com/evgenivanovi/gophkeeper/internal/client/usecase/config"
	authshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	authusershareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	authmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	commonmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
)

/* __________________________________________________ */

type SigninUsecase interface {
	Execute(context.Context, SignInRequest) (SignInResponse, error)
}

type SigninUsecaseService struct {
	authClient          auth.AuthAPI
	addUserUsecase      configuc.AddUserUsecase
	setContextUsecase   configuc.SetUserUsecase
	createConfigUsecase configuc.CreateConfigUsecase
	configManager       config.Manager
}

func ProvideSigninUsecaseService(
	authClient auth.AuthAPI,
	addUserUsecase configuc.AddUserUsecase,
	setContextUsecase configuc.SetUserUsecase,
	createConfigUsecase configuc.CreateConfigUsecase,
	configManager config.Manager,
) *SigninUsecaseService {
	return &SigninUsecaseService{
		authClient:          authClient,
		addUserUsecase:      addUserUsecase,
		setContextUsecase:   setContextUsecase,
		createConfigUsecase: createConfigUsecase,
		configManager:       configManager,
	}
}

func (uc *SigninUsecaseService) Execute(
	ctx context.Context,
	request SignInRequest,
) (SignInResponse, error) {

	response, err := uc.executeSignin(ctx, request)
	if err != nil {
		return response, err
	}

	err = uc.executeConfig(ctx, response)
	if err != nil {
		return response, err
	}

	return response, nil

}

func (uc *SigninUsecaseService) executeSignin(
	ctx context.Context,
	request SignInRequest,
) (response SignInResponse, err error) {

	credentials := authshareddm.NewCredentials(
		request.Payload.Credentials.Username,
		request.Payload.Credentials.Password,
	)

	usr, err := uc.authClient.Signin(ctx, credentials)

	if err != nil {
		return SignInResponse{}, err
	}

	return uc.toResponse(usr), nil

}

func (uc *SigninUsecaseService) executeConfig(
	ctx context.Context,
	response SignInResponse,
) error {

	err := uc.createConfigUsecase.Execute(ctx)
	if err != nil {
		return err
	}

	actions := make([]config.ConfigAction, 0)

	obj := ToUserObject(response.Payload.User, response.Payload.Session)
	actions = append(actions, uc.configManager.AddUserAction(ctx, obj))

	current := response.Payload.User.Data.Username
	actions = append(actions, uc.configManager.SetCurrentUserAction(ctx, current))

	options := common.MustOptionsFromCtx(ctx)
	return uc.configManager.Within(ctx, options.Config, actions...)

}

/* __________________________________________________ */

func (uc *SigninUsecaseService) toResponse(entity authusershareddm.AuthUser) SignInResponse {
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
				ID: entity.Identity().ID(),
				Data: authmd.UserDataModel{
					Username: entity.Data().Credentials.Username(),
				},
				Metadata: commonmd.FromMetadata(entity.Metadata()),
			},
		},
	}
}

/* __________________________________________________ */
