package secret

import (
	"context"

	secretdm "github.com/evgenivanovi/gophkeeper/internal/server/domain/secret"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	secretshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	commonsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
	secretmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type CreateDecodedSecretUsecase interface {
	Execute(
		ctx context.Context, request CreateDecodedSecretRequest,
	) (CreateEncodedSecretResponse, error)
}

type CreateDecodedSecretUsecaseService struct {
	manager secretdm.Manager
}

func ProvideCreateDecodedSecretUsecaseService(
	manager secretdm.Manager,
) *CreateDecodedSecretUsecaseService {
	return &CreateDecodedSecretUsecaseService{
		manager: manager,
	}
}

func (uc *CreateDecodedSecretUsecaseService) Execute(
	ctx context.Context, request CreateDecodedSecretRequest,
) (CreateEncodedSecretResponse, error) {

	data, err := uc.toDecodedSecretData(request)
	if err != nil {
		return CreateEncodedSecretResponse{}, err
	}

	response, err := uc.manager.CreateDecoded(ctx, data)
	if err != nil {
		return CreateEncodedSecretResponse{}, err
	}

	return uc.toEncodedSecretResponse(response), err

}

func (uc *CreateDecodedSecretUsecaseService) toDecodedSecretData(
	request CreateDecodedSecretRequest,
) (secretshareddm.OwnedDecodedSecretData, error) {

	user := common.NewUserID(
		request.Payload.Secret.UserID,
	)

	kind, err := secretshareddm.TypeFromString(
		request.Payload.Secret.Type,
	)
	if err != nil {
		return secretshareddm.OwnedDecodedSecretData{}, err
	}

	content := secretmd.ToSecretContent(
		request.Payload.Secret.Content,
	)

	data := secretshareddm.NewOwnedDecodedSecretData(
		user, request.Payload.Secret.Name, kind, content,
	)

	return *data, nil

}

func (uc *CreateDecodedSecretUsecaseService) toEncodedSecretResponse(
	entity secretshareddm.OwnedEncodedSecret,
) CreateEncodedSecretResponse {
	return CreateEncodedSecretResponse{
		Payload: CreateEncodedSecretResponsePayload{
			Secret: secretmd.OwnedEncodedSecretModel{
				ID:       entity.Identity().ID(),
				Data:     secretmd.FromOwnedEncodedSecretData(entity.Data()),
				Metadata: commonsharedmd.FromMetadata(entity.Metadata()),
			},
		},
	}
}

/* __________________________________________________ */
