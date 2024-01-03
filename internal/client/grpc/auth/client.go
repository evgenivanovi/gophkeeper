package auth

import (
	"context"

	authapi "github.com/evgenivanovi/gophkeeper/api/pb/auth"
	authdm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	userdm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gophkeeper/pkg/proto"
	errx "github.com/evgenivanovi/gpl/err"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* __________________________________________________ */

type APIService struct {
	client authapi.AuthAPIClient
}

func ProvideAPIService(client authapi.AuthAPIClient) *APIService {
	return &APIService{
		client: client,
	}
}

func (a *APIService) Signin(
	ctx context.Context, credentials authdm.Credentials,
) (userdm.AuthUser, error) {
	in := a.buildSigninRequest(credentials)

	out, err := a.client.Signin(ctx, in)
	if err != nil {
		return userdm.NewEmptyAuthUser(), a.translateSigninResponseError(err)
	}

	return a.buildSigninResponse(out, credentials), nil
}

func (a *APIService) buildSigninRequest(credentials authdm.Credentials) *authapi.SigninRequest {
	return &authapi.SigninRequest{
		Payload: &authapi.SigninRequest_SigninRequestPayload{
			Credentials: &authapi.Credentials{
				Username: credentials.Username(),
				Password: credentials.Password(),
			},
		},
	}
}

func (a *APIService) buildSigninResponse(
	out *authapi.SigninResponse, credentials authdm.Credentials,
) userdm.AuthUser {

	id := common.NewUserID(
		out.Payload.User.Id,
	)

	data := userdm.NewAuthUserData(
		session.NewSessionID(out.Payload.Session.Id),
		*token.NewTokens(
			*token.NewToken(
				out.Payload.Session.Tokens.Access.Token,
				out.Payload.Session.Tokens.Access.Expiration.AsTime(),
			),
			*token.NewToken(
				out.Payload.Session.Tokens.Refresh.Token,
				out.Payload.Session.Tokens.Refresh.Expiration.AsTime(),
			),
		),
		credentials,
	)

	metadata := core.NewMetadata(
		out.Payload.User.Metadata.CreatedAt.AsTime(),
		proto.NewTimeFromOptionalTimestamp(out.Payload.User.Metadata.UpdatedAt),
		proto.NewTimeFromOptionalTimestamp(out.Payload.User.Metadata.DeletedAt),
	)

	return *userdm.NewAuthUser(id, *data, *metadata)

}

func (a *APIService) translateSigninResponseError(err error) error {

	code, _ := status.FromError(err)

	if code.Code() == codes.NotFound {
		return errx.NewErrorWithEntityCode(
			userdm.ErrorUserEntity, core.ErrorNotFoundCode,
		)
	}

	if code.Code() == codes.Unauthenticated {
		return errx.NewErrorWithEntityCode(
			userdm.ErrorUserEntity, core.ErrorUnauthenticatedCode,
		)
	}

	if code.Code() == codes.Internal {
		return errx.NewErrorWithEntityCode(
			userdm.ErrorUserEntity, errx.ErrorInternalMessage,
		)
	}

	return errx.NewErrorWithEntityCode(
		userdm.ErrorUserEntity, errx.ErrorInternalMessage,
	)

}

func (a *APIService) Signup(
	ctx context.Context, credentials authdm.Credentials,
) (userdm.User, error) {
	in := a.buildSignupRequest(credentials)

	out, err := a.client.Signup(ctx, in)
	if err != nil {
		return userdm.NewEmptyUser(), a.translateSignupResponseError(err)
	}

	return a.buildSignupResponse(out, credentials), nil
}

func (a *APIService) buildSignupRequest(
	credentials authdm.Credentials,
) *authapi.SignupRequest {
	return &authapi.SignupRequest{
		Payload: &authapi.SignupRequest_SignupRequestPayload{
			Credentials: &authapi.Credentials{
				Username: credentials.Username(),
				Password: credentials.Password(),
			},
		},
	}
}

func (a *APIService) buildSignupResponse(
	out *authapi.SignupResponse, credentials authdm.Credentials,
) userdm.User {

	id := common.NewUserID(
		out.Payload.User.Id,
	)

	data := userdm.NewUserData(
		credentials,
	)

	metadata := core.NewMetadata(
		out.Payload.User.Metadata.CreatedAt.AsTime(),
		proto.NewTimeFromOptionalTimestamp(out.Payload.User.Metadata.UpdatedAt),
		proto.NewTimeFromOptionalTimestamp(out.Payload.User.Metadata.DeletedAt),
	)

	return *userdm.NewUser(id, *data, *metadata)

}

func (a *APIService) translateSignupResponseError(err error) error {

	code, _ := status.FromError(err)

	if code.Code() == codes.AlreadyExists {
		return errx.NewErrorWithEntityCode(
			userdm.ErrorUserEntity, core.ErrorExistsCode,
		)
	}

	if code.Code() == codes.Internal {
		return errx.NewErrorWithEntityCode(
			userdm.ErrorUserEntity, errx.ErrorInternalMessage,
		)
	}

	return errx.NewErrorWithEntityCode(
		userdm.ErrorUserEntity, errx.ErrorInternalMessage,
	)

}

func (a *APIService) SignupAndSignin(
	ctx context.Context, credentials authdm.Credentials,
) (userdm.AuthUser, error) {

	_, err := a.Signup(ctx, credentials)
	if err != nil {
		return userdm.NewEmptyAuthUser(), err
	}

	usr, err := a.Signin(ctx, credentials)
	if err != nil {
		return userdm.NewEmptyAuthUser(), err
	}

	return usr, err

}

/* __________________________________________________ */
