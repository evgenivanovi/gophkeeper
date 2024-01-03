package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/domain/secret"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	secretshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type GetEncodedByNameUsecase interface {
	Execute(
		ctx context.Context, user common.UserID, name string,
	) (secretsharedmd.OwnedEncodedSecretModel, error)
}

type GetEncodedByNameUsecaseService struct {
	seeker secret.Seeker
	encdec secretshareddm.OwnedSecretEncoderDecoder
}

func ProvideGetEncodedByNameUsecaseService(
	seeker secret.Seeker,
	encdec secretshareddm.OwnedSecretEncoderDecoder,
) *GetEncodedByNameUsecaseService {
	return &GetEncodedByNameUsecaseService{
		seeker: seeker,
		encdec: encdec,
	}
}

func (uc *GetEncodedByNameUsecaseService) Execute(
	ctx context.Context, user common.UserID, name string,
) (secretsharedmd.OwnedEncodedSecretModel, error) {

	sec, err := uc.seeker.GetByName(ctx, user, name)
	if err != nil {
		return secretsharedmd.OwnedEncodedSecretModel{}, err
	}

	return secretsharedmd.FromOwnedEncodedSecret(sec), nil

}

/* __________________________________________________ */
