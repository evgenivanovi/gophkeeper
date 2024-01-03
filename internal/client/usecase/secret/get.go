package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/secret"
	secredshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type GetSecretUsecase interface {
	Execute(
		ctx context.Context, name string,
	) (secretsharedmd.DecodedSecretModel, error)
}

type GetSecretUsecaseService struct {
	secretAPI     secret.SeekerAPI
	configManager config.Manager
	encdec        secredshareddm.SecretEncoderDecoder
}

func ProvideGetSecretUsecaseService(
	secretAPI secret.SeekerAPI,
	configManager config.Manager,
	encdec secredshareddm.SecretEncoderDecoder,
) *GetSecretUsecaseService {
	return &GetSecretUsecaseService{
		secretAPI:     secretAPI,
		configManager: configManager,
		encdec:        encdec,
	}
}

func (uc *GetSecretUsecaseService) Execute(
	ctx context.Context, name string,
) (secretsharedmd.DecodedSecretModel, error) {

	user, err := uc.configManager.GetCurrentUser(
		ctx, common.MustOptionsFromCtx(ctx).Config,
	)

	if err != nil {
		return secretsharedmd.DecodedSecretModel{}, err
	}

	sec, err := uc.secretAPI.GetByName(ctx, user, name)

	if err != nil {
		return secretsharedmd.DecodedSecretModel{}, err
	}

	decoded, err := uc.encdec.Decode(sec)

	if err != nil {
		return secretsharedmd.DecodedSecretModel{}, err
	}

	return secretsharedmd.FromDecodedSecret(decoded), nil

}

/* __________________________________________________ */
