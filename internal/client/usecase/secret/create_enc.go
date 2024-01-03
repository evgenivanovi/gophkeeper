package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	secretdm "github.com/evgenivanovi/gophkeeper/internal/client/domain/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type CreateEncodedSecretUsecase interface {
	Execute(
		ctx context.Context, data secretsharedmd.EncodedSecretDataModel,
	) error
}

type CreateEncodedSecretUsecaseService struct {
	secretAPI     secretdm.SecretManagementAPI
	configManager config.Manager
}

func ProvideCreateEncodedSecretUsecaseService(
	secretAPI secretdm.SecretManagementAPI,
	configManager config.Manager,
) *CreateEncodedSecretUsecaseService {
	return &CreateEncodedSecretUsecaseService{
		secretAPI:     secretAPI,
		configManager: configManager,
	}
}

func (uc *CreateEncodedSecretUsecaseService) Execute(
	ctx context.Context, data secretsharedmd.EncodedSecretDataModel,
) error {

	user, err := uc.configManager.GetCurrentUser(
		ctx, common.MustOptionsFromCtx(ctx).Config,
	)

	if err != nil {
		return err
	}

	_, err = uc.secretAPI.CreateEncoded(
		ctx, user, secretsharedmd.ToEncodedSecretData(data),
	)

	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
