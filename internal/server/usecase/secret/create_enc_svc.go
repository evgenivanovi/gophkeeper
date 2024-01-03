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

type CreateEncodedSecretUsecase interface {
	Execute(
		ctx context.Context, request CreateEncodedSecretRequest,
	) (CreateEncodedSecretResponse, error)
}

type CreateEncodedSecretUsecaseService struct {
	manager secretdm.Manager
}

func ProvideCreateEncodedSecretUsecaseService(
	manager secretdm.Manager,
) *CreateEncodedSecretUsecaseService {
	return &CreateEncodedSecretUsecaseService{
		manager: manager,
	}
}

func (uc *CreateEncodedSecretUsecaseService) Execute(
	ctx context.Context, request CreateEncodedSecretRequest,
) (CreateEncodedSecretResponse, error) {

	data, err := uc.toEncodedSecretData(request)
	if err != nil {
		return CreateEncodedSecretResponse{}, err
	}

	response, err := uc.manager.CreateEncoded(ctx, data)
	if err != nil {
		return CreateEncodedSecretResponse{}, err
	}

	return uc.toEncodedSecretResponse(response), err

}

func (uc *CreateEncodedSecretUsecaseService) toEncodedSecretData(
	request CreateEncodedSecretRequest,
) (secretshareddm.OwnedEncodedSecretData, error) {

	user := common.NewUserID(
		request.Payload.Secret.UserID,
	)

	kind, err := secretshareddm.TypeFromString(
		request.Payload.Secret.Type,
	)
	if err != nil {
		return secretshareddm.OwnedEncodedSecretData{}, err
	}

	data := secretshareddm.NewOwnedEncodedSecretData(
		user, request.Payload.Secret.Name, kind, request.Payload.Secret.Content,
	)

	return *data, nil

}

func (uc *CreateEncodedSecretUsecaseService) toEncodedSecretResponse(
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
