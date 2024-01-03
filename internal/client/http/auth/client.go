package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/evgenivanovi/gophkeeper/api/http/authapi"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	errx "github.com/evgenivanovi/gpl/err"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/go-resty/resty/v2"
)

/* __________________________________________________ */

type APIService struct {
	client *resty.Client
}

func ProvideAPIService(client *resty.Client) *APIService {
	return &APIService{
		client: client,
	}
}

func (svc *APIService) Signin(
	ctx context.Context, credentials auth.Credentials,
) (user.AuthUser, error) {
	resp, err := svc.executeSignin(ctx, credentials)
	return svc.buildSigninResponse(resp, err, credentials)
}

func (svc *APIService) executeSignin(
	ctx context.Context, credentials auth.Credentials,
) (*resty.Response, error) {

	req := authapi.SigninRequest{
		Username: credentials.Username(),
		Password: credentials.Password(),
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return svc.client.
		R().
		SetContext(ctx).
		SetHeader(
			headers.ContentTypeKey.String(),
			headers.TypeApplicationJSON.String(),
		).
		SetHeader(
			headers.AcceptKey.String(),
			headers.TypeApplicationJSON.String(),
		).
		SetBody(body).
		SetResult(&authapi.SigninResponse{}).
		Post(authapi.SigninEndpoint.String())

}

func (svc *APIService) buildSigninResponse(
	resp *resty.Response, err error, credentials auth.Credentials,
) (user.AuthUser, error) {

	if err != nil {
		return user.NewEmptyAuthUser(), err
	}

	status := resp.StatusCode()

	if status == http.StatusOK {
		response := resp.Result().(*authapi.SigninResponse)
		return svc.mapSigninResponse(*response, credentials), nil
	}

	if status == http.StatusNotFound {
		return user.NewEmptyAuthUser(), errx.NewErrorWithEntityCode(
			user.ErrorUserEntity, core.ErrorNotFoundCode,
		)
	}

	if status == http.StatusUnauthorized {
		return user.NewEmptyAuthUser(), errx.NewErrorWithEntityCode(
			user.ErrorUserEntity, core.ErrorUnauthenticatedCode,
		)
	}

	return user.NewEmptyAuthUser(), errx.NewErrorWithEntityCode(
		user.ErrorUserEntity, errx.ErrorInternalMessage,
	)

}

func (svc *APIService) mapSigninResponse(
	response authapi.SigninResponse, credentials auth.Credentials,
) user.AuthUser {

	id := common.NewUserID(response.Payload.User.ID)

	sessionID := session.NewSessionID(response.Payload.Session.ID)

	tokens := *token.NewTokens(
		*token.NewToken(
			response.Payload.Session.Tokens.AccessToken.Token,
			response.Payload.Session.Tokens.AccessToken.ExpiresAt,
		),
		*token.NewToken(
			response.Payload.Session.Tokens.RefreshToken.Token,
			response.Payload.Session.Tokens.RefreshToken.ExpiresAt,
		),
	)

	data := *user.NewAuthUserData(
		sessionID, tokens, credentials,
	)

	metadata := *core.NewMetadata(
		response.Payload.User.Metadata.CreatedAt,
		response.Payload.User.Metadata.UpdatedAt,
		response.Payload.User.Metadata.DeletedAt,
	)

	return *user.NewAuthUser(id, data, metadata)

}

func (svc *APIService) Signup(
	ctx context.Context, credentials auth.Credentials,
) (user.User, error) {
	resp, err := svc.executeSignup(ctx, credentials)
	return svc.buildSignupResponse(resp, err, credentials)
}

func (svc *APIService) executeSignup(
	ctx context.Context, credentials auth.Credentials,
) (*resty.Response, error) {

	req := authapi.SignupRequest{
		Username: credentials.Username(),
		Password: credentials.Password(),
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return svc.client.
		R().
		SetContext(ctx).
		SetHeader(
			headers.ContentTypeKey.String(),
			headers.TypeApplicationJSON.String(),
		).
		SetHeader(
			headers.AcceptKey.String(),
			headers.TypeApplicationJSON.String(),
		).
		SetBody(body).
		SetResult(&authapi.SignupResponse{}).
		Post(authapi.SignupEndpoint.String())

}

func (svc *APIService) buildSignupResponse(
	resp *resty.Response, err error, credentials auth.Credentials,
) (user.User, error) {

	if err != nil {
		return user.NewEmptyUser(), err
	}

	status := resp.StatusCode()

	if status == http.StatusOK {
		response := resp.Result().(*authapi.SignupResponse)
		return svc.mapSignupResponse(*response, credentials), nil
	}

	if status == http.StatusConflict {
		return user.NewEmptyUser(), errx.NewErrorWithEntityCode(
			user.ErrorUserEntity, core.ErrorExistsCode,
		)
	}

	return user.NewEmptyUser(), errx.NewErrorWithEntityCode(
		user.ErrorUserEntity, errx.ErrorInternalMessage,
	)

}

func (svc *APIService) mapSignupResponse(
	response authapi.SignupResponse, credentials auth.Credentials,
) user.User {

	id := common.NewUserID(response.Payload.User.ID)

	data := *user.NewUserData(credentials)

	metadata := *core.NewMetadata(
		response.Payload.User.Metadata.CreatedAt,
		response.Payload.User.Metadata.UpdatedAt,
		response.Payload.User.Metadata.DeletedAt,
	)

	return *user.NewUser(id, data, metadata)

}

func (svc *APIService) SignupAndSignin(
	ctx context.Context, credentials auth.Credentials,
) (user.AuthUser, error) {

	_, err := svc.Signup(ctx, credentials)
	if err != nil {
		return user.NewEmptyAuthUser(), err
	}

	usr, err := svc.Signin(ctx, credentials)
	if err != nil {
		return user.NewEmptyAuthUser(), err
	}

	return usr, err

}

/* __________________________________________________ */
