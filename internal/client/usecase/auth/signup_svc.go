package auth

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	configuc "github.com/evgenivanovi/gophkeeper/internal/client/usecase/config"
	authshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	userauthshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	authmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	commonmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
)

/* __________________________________________________ */

type SignupUsecase interface {
	Execute(context.Context, SignUpRequest) (SignUpResponse, error)
}

type SignupUsecaseService struct {
	authClient          auth.AuthAPI
	addUserUsecase      configuc.AddUserUsecase
	setContextUsecase   configuc.SetUserUsecase
	createConfigUsecase configuc.CreateConfigUsecase
	configManager       config.Manager
}

func ProvideSignupUsecaseService(
	authClient auth.AuthAPI,
	addUserUsecase configuc.AddUserUsecase,
	setContextUsecase configuc.SetUserUsecase,
	createConfigUsecase configuc.CreateConfigUsecase,
	configManager config.Manager,
) *SignupUsecaseService {
	return &SignupUsecaseService{
		authClient:          authClient,
		addUserUsecase:      addUserUsecase,
		setContextUsecase:   setContextUsecase,
		createConfigUsecase: createConfigUsecase,
		configManager:       configManager,
	}
}

func (uc *SignupUsecaseService) Execute(
	ctx context.Context,
	request SignUpRequest,
) (SignUpResponse, error) {

	response, err := uc.executeSignup(ctx, request)
	if err != nil {
		return response, err
	}

	err = uc.executeConfig(ctx, response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (uc *SignupUsecaseService) executeSignup(
	ctx context.Context,
	request SignUpRequest,
) (SignUpResponse, error) {

	credentials := authshareddm.NewCredentials(
		request.Payload.Credentials.Username,
		request.Payload.Credentials.Password,
	)

	usr, err := uc.authClient.SignupAndSignin(ctx, credentials)

	if err != nil {
		return SignUpResponse{}, err
	}

	return uc.toResponse(usr), nil

}

func (uc *SignupUsecaseService) executeConfig(
	ctx context.Context,
	response SignUpResponse,
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

	path := common.MustOptionsFromCtx(ctx).Config
	return uc.configManager.Within(ctx, path, actions...)

}

/* __________________________________________________ */

func (uc *SignupUsecaseService) toResponse(entity userauthshareddm.AuthUser) SignUpResponse {
	return SignUpResponse{
		Payload: SignUpResponsePayload{
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
