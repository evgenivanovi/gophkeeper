package secret

import (
	"context"

	secretapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/private"
	secretuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/secret"
	"google.golang.org/grpc"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type InternalSecretManagementAPI struct {
	secretapi.UnimplementedInternalSecretManagementAPIServer

	createDecoded secretuc.CreateDecodedSecretUsecase
	createEncoded secretuc.CreateEncodedSecretUsecase
}

func ProvideInternalSecretManagementAPI(
	createDecoded secretuc.CreateDecodedSecretUsecase,
	createEncoded secretuc.CreateEncodedSecretUsecase,
) *InternalSecretManagementAPI {
	return &InternalSecretManagementAPI{
		createDecoded: createDecoded,
		createEncoded: createEncoded,
	}
}

func (a *InternalSecretManagementAPI) RegisterService(registrar grpc.ServiceRegistrar) {
	secretapi.RegisterInternalSecretManagementAPIServer(registrar, a)
}

func (a *InternalSecretManagementAPI) CreateDecoded(
	ctx context.Context, in *secretapi.CreateDecodedSecretRequest,
) (*secretapi.CreateDecodedSecretResponse, error) {

	request := secretuc.CreateDecodedSecretRequest{
		Payload: secretuc.CreateDecodedSecretRequestPayload{
			Secret: ToOwnedDecodedSecretDataModel(in.GetPayload().GetData(), in.GetPayload().GetUserId()),
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

func (a *InternalSecretManagementAPI) CreateEncoded(
	ctx context.Context, in *secretapi.CreateEncodedSecretRequest,
) (*secretapi.CreateEncodedSecretResponse, error) {

	request := secretuc.CreateEncodedSecretRequest{
		Payload: secretuc.CreateEncodedSecretRequestPayload{
			Secret: ToOwnedEncodedSecretDataModel(in.GetPayload().GetData(), in.GetPayload().GetUserId()),
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
