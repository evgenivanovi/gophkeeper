package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/postgres"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/pg"
)

/* __________________________________________________ */

type PGWriteRepositoryService struct {
	requester pg.WriteRequester
}

func ProvidePGWriteRepositoryService(
	requester pg.WriteRequester,
) *PGWriteRepositoryService {
	return &PGWriteRepositoryService{
		requester: requester,
	}
}

func (r *PGWriteRepositoryService) AutoSave(
	ctx context.Context, data secret.OwnedEncodedSecretData, metadata core.Metadata,
) (*secret.OwnedEncodedSecret, error) {

	var id int64
	command, args := insertOneStatement(FromSecretData(data, metadata))

	err := r.requester.ExecReturningWithDefaultOnError(ctx, &id, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return secret.NewOwnedEncodedSecret(secret.NewSecretID(id), data, metadata), nil

}

func (r *PGWriteRepositoryService) translateError(err error) error {

	if err == nil {
		return nil
	}

	err = pg.WithEntity(err, secret.ErrorSecretEntity)
	err = postgres.TranslateWriteError(err)

	if pg.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil
	}

	return err

}

/* __________________________________________________ */
