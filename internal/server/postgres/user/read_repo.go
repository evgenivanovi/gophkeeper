package user

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/model"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/table"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	errx "github.com/evgenivanovi/gpl/err"
	"github.com/evgenivanovi/gpl/pg"
	"github.com/evgenivanovi/gpl/search"
	pgsearch "github.com/evgenivanovi/gpl/search/jet/pg"
	pgjet "github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

type PGReadRepositoryService struct {
	requester     pg.ReadRequester
	searchMapping map[search.Key]pgjet.Column
	orderMapping  map[search.Key]pgjet.Column
}

func ProvidePGReadRepositoryService(
	requester pg.ReadRequester,
) *PGReadRepositoryService {

	searchMapping := make(map[search.Key]pgjet.Column)
	searchMapping[auth.UserIDSearchKey] = table.Users.ID
	searchMapping[auth.UsernameSearchKey] = table.Users.Username

	orderMapping := make(map[search.Key]pgjet.Column)

	return &PGReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *PGReadRepositoryService) GetByID(
	ctx context.Context, id common.UserID,
) (*user.User, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			user.ErrorUserEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *PGReadRepositoryService) FindByID(
	ctx context.Context, id common.UserID,
) (*user.User, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(auth.UserIDCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *PGReadRepositoryService) FindByIDs(
	ctx context.Context, ids []common.UserID,
) ([]*user.User, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(auth.UserIDsCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *PGReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*user.User, error) {
	var dst model.Users
	query, args := r.query(spec)

	err := r.requester.ExecOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUser(&dst), nil
}

func (r *PGReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*user.User, error) {
	var dst = make([]*model.Users, 0)
	query, args := r.query(spec)

	err := r.requester.ExecManyWithScan(ctx, scanManyFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToUsers(dst), nil
}

func (r *PGReadRepositoryService) query(
	spec search.Specification,
) (string, []interface{}) {
	searchExp := pgsearch.SearchExpression(spec, r.searchMapping)
	orderExp := pgsearch.OrderExpression(spec, r.orderMapping)
	return buildQuery(searchExp, orderExp, nil, *spec.SliceConditions())
}

func (r *PGReadRepositoryService) translateError(err error) error {

	if err == nil {
		return nil
	}

	err = pg.WithEntity(err, user.ErrorUserEntity)
	err = postgres.TranslateReadError(err)

	if pg.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil
	}

	return err

}

/* __________________________________________________ */
