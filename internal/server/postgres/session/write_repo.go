package session

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/postgres"
	session "github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
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

func (r *PGWriteRepositoryService) NonAutoSave(
	ctx context.Context, data session.Session,
) (*session.Session, error) {

	command, args := insertOneStatement(FromSession(data))

	err := r.requester.ExecWithDefaultOnError(ctx, command, args...)
	err = r.translateError(err)

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *PGWriteRepositoryService) NonAutoSaveAll(
	ctx context.Context, data []session.Session,
) error {
	command, args := insertAllStatement(FromSessions(data))
	err := r.requester.ExecWithDefaultOnError(ctx, command, args...)
	return r.translateError(err)
}

func (r *PGWriteRepositoryService) translateError(err error) error {

	if err == nil {
		return nil
	}

	err = pg.WithEntity(err, session.ErrorSessionEntity)
	err = postgres.TranslateWriteError(err)

	if pg.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil
	}

	return err

}

/* __________________________________________________ */
