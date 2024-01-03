package secret

import (
	"context"

	secretapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/public"
	secretuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/secret"
	"github.com/evgenivanovi/gophkeeper/internal/shared/util/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretManagementAPI struct {
	secretapi.UnimplementedSecretManagementAPIServer

	createDecoded secretuc.CreateDecodedSecretUsecase
	createEncoded secretuc.CreateEncodedSecretUsecase
}

func ProvideSecretManagementAPI(
	createDecoded secretuc.CreateDecodedSecretUsecase,
	createEncoded secretuc.CreateEncodedSecretUsecase,
) *SecretManagementAPI {
	return &SecretManagementAPI{
		createDecoded: createDecoded,
		createEncoded: createEncoded,
	}
}

func (a *SecretManagementAPI) RegisterService(registrar grpc.ServiceRegistrar) {
	secretapi.RegisterSecretManagementAPIServer(registrar, a)
}

func (a *SecretManagementAPI) CreateDecoded(
	ctx context.Context, in *secretapi.CreateDecodedSecretRequest,
) (*secretapi.CreateDecodedSecretResponse, error) {

	user := auth.FromCtx(ctx)
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "user not defined")
	}

	request := secretuc.CreateDecodedSecretRequest{
		Payload: secretuc.CreateDecodedSecretRequestPayload{
			Secret: ToOwnedDecodedSecretDataModel(in.GetPayload().GetData(), user.UserID),
		},
	}

	response, err := a.createDecoded.Execute(ctx, request)
	if err != nil {
		return &secretapi.CreateDecodedSecretResponse{}, err
	}

	out := secretapi.CreateDecodedSecretResponse{
		Payload: &secretapi.CreateDecodedSecretResponse_Payload{
			Data: FromOwnedEncodedSecretModel(response.Payload.Secret),
		},
	}

	return &out, nil

}

func (a *SecretManagementAPI) CreateEncoded(
	ctx context.Context, in *secretapi.CreateEncodedSecretRequest,
) (*secretapi.CreateEncodedSecretResponse, error) {

	user := auth.FromCtx(ctx)
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "user not defined")
	}

	request := secretuc.CreateEncodedSecretRequest{
		Payload: secretuc.CreateEncodedSecretRequestPayload{
			Secret: ToOwnedEncodedSecretDataModel(in.GetPayload().GetData(), user.UserID),
		},
	}

	response, err := a.createEncoded.Execute(ctx, request)
	if err != nil {
		return &secretapi.CreateEncodedSecretResponse{}, err
	}

	out := secretapi.CreateEncodedSecretResponse{
		Payload: &secretapi.CreateEncodedSecretResponse_Payload{
			Data: FromOwnedEncodedSecretModel(response.Payload.Secret),
		},
	}

	return &out, nil

}

/* __________________________________________________ */
