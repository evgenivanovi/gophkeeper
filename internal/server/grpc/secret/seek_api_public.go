package secret

import (
	"context"

	secretapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/public"
	"github.com/evgenivanovi/gophkeeper/internal/server/usecase/secret"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gophkeeper/internal/shared/util/auth"
	errx "github.com/evgenivanovi/gpl/err"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretSeekingAPI struct {
	secretapi.UnimplementedSecretSeekingAPIServer

	get secret.GetEncodedByNameUsecase
}

func ProvideSecretSeekingAPI(
	get secret.GetEncodedByNameUsecase,
) *SecretSeekingAPI {
	return &SecretSeekingAPI{
		get: get,
	}
}

func (a *SecretSeekingAPI) RegisterService(registrar grpc.ServiceRegistrar) {
	secretapi.RegisterSecretSeekingAPIServer(registrar, a)
}

func (a *SecretSeekingAPI) GetByName(
	ctx context.Context, in *secretapi.GetByNameSecretRequest,
) (*secretapi.GetSecretResponse, error) {

	user := auth.FromCtx(ctx)
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "user not defined")
	}

	response, err := a.get.Execute(
		ctx, common.NewUserID(user.UserID), in.GetPayload().GetName(),
	)

	if err != nil {
		return nil, a.translateResponseError(err)
	}

	out := &secretapi.GetSecretResponse{
		Payload: &secretapi.GetSecretResponse_Payload{
			Data: FromOwnedEncodedSecretModel(response),
		},
	}

	return out, nil

}

func (a *SecretSeekingAPI) translateResponseError(err error) error {

	code := errx.ErrorCode(err)

	if code == core.ErrorNotFoundCode {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())

}

/* __________________________________________________ */
