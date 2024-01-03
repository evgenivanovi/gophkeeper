package auth

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/api/pb/auth"
	"github.com/evgenivanovi/gophkeeper/api/pb/common"
	authuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	authmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	"github.com/evgenivanovi/gophkeeper/pkg/proto"
	errx "github.com/evgenivanovi/gpl/err"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type AuthAPI struct {
	auth.UnimplementedAuthAPIServer

	signin authuc.SigninUsecase
	signup authuc.SignupUsecase
}

func ProvideAuthAPI(
	signin authuc.SigninUsecase,
	signup authuc.SignupUsecase,
) *AuthAPI {
	return &AuthAPI{
		signin: signin,
		signup: signup,
	}
}

func (a *AuthAPI) RegisterService(registrar grpc.ServiceRegistrar) {
	auth.RegisterAuthAPIServer(registrar, a)
}

func (a *AuthAPI) Signin(
	ctx context.Context, in *auth.SigninRequest,
) (*auth.SigninResponse, error) {

	request := authuc.SignInRequest{
		Payload: authuc.SignInRequestPayload{
			Credentials: authmd.CredentialsModel{
				Username: in.Payload.Credentials.Username,
				Password: in.Payload.Credentials.Password,
			},
		},
	}

	response, err := a.signin.Execute(ctx, request)
	if err != nil {
		return nil, a.translateSigninResponseError(err)
	}

	out := auth.SigninResponse{
		Payload: &auth.SigninResponse_SigninResponsePayload{
			Session: &auth.Session{
				Id: response.Payload.Session.ID,
				Tokens: &auth.Tokens{
					Access: &auth.Token{
						Token:      response.Payload.Session.Tokens.AccessToken.Token,
						Expiration: timestamppb.New(response.Payload.Session.Tokens.AccessToken.ExpiresAt),
					},
					Refresh: &auth.Token{
						Token:      response.Payload.Session.Tokens.RefreshToken.Token,
						Expiration: timestamppb.New(response.Payload.Session.Tokens.RefreshToken.ExpiresAt),
					},
				},
			},
			User: &auth.User{
				Id: response.Payload.User.ID,
				Metadata: &common.Metadata{
					CreatedAt: timestamppb.New(response.Payload.User.Metadata.CreatedAt),
					UpdatedAt: proto.NewOptionalTimestampFromTime(response.Payload.User.Metadata.UpdatedAt),
					DeletedAt: proto.NewOptionalTimestampFromTime(response.Payload.User.Metadata.DeletedAt),
				},
			},
		},
	}

	return &out, nil

}

func (a *AuthAPI) translateSigninResponseError(err error) error {

	code := errx.ErrorCode(err)
	msg := errx.ErrorMessage(err)

	if code == core.ErrorNotFoundCode {
		return status.Error(codes.NotFound, msg)
	}

	if code == core.ErrorUnauthenticatedCode {
		return status.Error(codes.Unauthenticated, msg)
	}

	if code == errx.ErrorInternalCode {
		return status.Error(codes.Internal, msg)
	}

	return status.Error(codes.Unknown, msg)

}

func (a *AuthAPI) Signup(
	ctx context.Context, in *auth.SignupRequest,
) (*auth.SignupResponse, error) {

	request := authuc.SignUpRequest{
		Payload: authuc.SignUpRequestPayload{
			Credentials: authmd.CredentialsModel{
				Username: in.Payload.Credentials.Username,
				Password: in.Payload.Credentials.Password,
			},
		},
	}

	response, err := a.signup.Execute(ctx, request)
	if err != nil {
		return nil, a.translateSignupResponseError(err)
	}

	out := auth.SignupResponse{
		Payload: &auth.SignupResponse_SignupResponsePayload{
			Session: &auth.Session{
				Id: response.Payload.Session.ID,
				Tokens: &auth.Tokens{
					Access: &auth.Token{
						Token:      response.Payload.Session.Tokens.AccessToken.Token,
						Expiration: timestamppb.New(response.Payload.Session.Tokens.AccessToken.ExpiresAt),
					},
					Refresh: &auth.Token{
						Token:      response.Payload.Session.Tokens.RefreshToken.Token,
						Expiration: timestamppb.New(response.Payload.Session.Tokens.RefreshToken.ExpiresAt),
					},
				},
			},
			User: &auth.User{
				Id: response.Payload.User.ID,
				Metadata: &common.Metadata{
					CreatedAt: timestamppb.New(response.Payload.User.Metadata.CreatedAt),
					UpdatedAt: proto.NewOptionalTimestampFromTime(response.Payload.User.Metadata.UpdatedAt),
					DeletedAt: proto.NewOptionalTimestampFromTime(response.Payload.User.Metadata.DeletedAt),
				},
			},
		},
	}

	return &out, nil

}

func (a *AuthAPI) translateSignupResponseError(err error) error {

	code := errx.ErrorCode(err)
	msg := errx.ErrorMessage(err)

	if code == core.ErrorExistsCode {
		return status.Error(codes.AlreadyExists, msg)
	}

	if code == errx.ErrorInternalCode {
		return status.Error(codes.Internal, msg)
	}

	return status.Error(codes.Unknown, msg)

}

/* __________________________________________________ */
