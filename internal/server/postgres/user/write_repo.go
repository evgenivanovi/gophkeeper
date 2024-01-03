package user

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/postgres"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gpl/pg"
	"github.com/evgenivanovi/gpl/std"
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
	ctx context.Context, data user.UserData, metadata core.Metadata,
) (*user.User, error) {

	var id int64
	command, args := insertOneStatement(FromUserData(data, metadata))

	err := r.requester.ExecReturningWithDefaultOnError(ctx, &id, command, args...)
	err = r.translateWriteError(err)

	if err != nil {
		return nil, err
	}

	return user.NewUser(common.NewUserID(id), data, metadata), nil

}

func (r *PGWriteRepositoryService) AutoSaveAll(
	ctx context.Context, datas []std.Pair[user.UserData, core.Metadata],
) error {
	command, args := insertAllStatement(FromUsersData(datas))
	err := r.requester.ExecWithDefaultOnError(ctx, command, args)
	return r.translateWriteError(err)
}

func (r *PGWriteRepositoryService) translateWriteError(err error) error {

	if err == nil {
		return nil
	}

	err = pg.WithEntity(err, user.ErrorUserEntity)
	err = postgres.TranslateWriteError(err)

	if pg.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil
	}

	return err

}

/* __________________________________________________ */
