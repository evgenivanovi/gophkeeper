package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/domain/secret"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	secretshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type GetDecodedByNameUsecase interface {
	Execute(
		ctx context.Context, user common.UserID, name string,
	) (secretsharedmd.OwnedDecodedSecretModel, error)
}

type GetDecodedByNameUsecaseService struct {
	seeker secret.Seeker
	encdec secretshareddm.OwnedSecretEncoderDecoder
}

func ProvideGetDecodedByNameUsecaseService(
	seeker secret.Seeker,
	encdec secretshareddm.OwnedSecretEncoderDecoder,
) *GetDecodedByNameUsecaseService {
	return &GetDecodedByNameUsecaseService{
		seeker: seeker,
		encdec: encdec,
	}
}

func (uc *GetDecodedByNameUsecaseService) Execute(
	ctx context.Context, user common.UserID, name string,
) (secretsharedmd.OwnedDecodedSecretModel, error) {

	sec, err := uc.seeker.GetByName(ctx, user, name)
	if err != nil {
		return secretsharedmd.OwnedDecodedSecretModel{}, err
	}

	decoded, err := uc.encdec.Decode(sec)
	if err != nil {
		return secretsharedmd.OwnedDecodedSecretModel{}, err
	}

	return secretsharedmd.FromOwnedDecodedSecret(decoded), nil

}

/* __________________________________________________ */
