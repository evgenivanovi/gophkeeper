package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	secretdm "github.com/evgenivanovi/gophkeeper/internal/client/domain/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type CreateDecodedSecretUsecase interface {
	Execute(
		ctx context.Context, data secretsharedmd.DecodedSecretDataModel,
	) error
}

type CreateDecodedSecretUsecaseService struct {
	secretAPI     secretdm.SecretManagementAPI
	configManager config.Manager
}

func ProvideCreateDecodedSecretUsecaseService(
	secretAPI secretdm.SecretManagementAPI,
	configManager config.Manager,
) *CreateDecodedSecretUsecaseService {
	return &CreateDecodedSecretUsecaseService{
		secretAPI:     secretAPI,
		configManager: configManager,
	}
}

func (uc *CreateDecodedSecretUsecaseService) Execute(
	ctx context.Context, data secretsharedmd.DecodedSecretDataModel,
) error {

	user, err := uc.configManager.GetCurrentUser(
		ctx, common.MustOptionsFromCtx(ctx).Config,
	)

	if err != nil {
		return err
	}

	_, err = uc.secretAPI.CreateDecoded(
		ctx, user, secretsharedmd.ToDecodedSecretData(data),
	)

	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
