package session

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/domain/auth"
	sessiondm "github.com/evgenivanovi/gophkeeper/internal/server/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/server/domain/secret"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/model"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/table"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
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
	searchMapping[secret.IDSearchKey] = table.Secrets.ID
	searchMapping[secret.NameSearchKey] = table.Secrets.Name
	searchMapping[auth.UserIDSearchKey] = table.Secrets.UserID

	orderMapping := make(map[search.Key]pgjet.Column)

	return &PGReadRepositoryService{
		requester:     requester,
		searchMapping: searchMapping,
		orderMapping:  orderMapping,
	}

}

func (r *PGReadRepositoryService) GetByID(
	ctx context.Context, id session.SessionID,
) (*session.Session, error) {

	res, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errx.NewErrorWithEntityCode(
			session.ErrorSessionEntity,
			core.ErrorNotFoundCode,
		)
	}

	return res, nil

}

func (r *PGReadRepositoryService) FindByID(
	ctx context.Context, id session.SessionID,
) (*session.Session, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(sessiondm.IdentityCondition(id))
	return r.FindOneBySpec(ctx, spec)
}

func (r *PGReadRepositoryService) FindByIDs(
	ctx context.Context, ids []session.SessionID,
) ([]*session.Session, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(sessiondm.IdentitiesCondition(ids))
	return r.FindManyBySpec(ctx, spec)
}

func (r *PGReadRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*session.Session, error) {
	var dst model.Sessions
	query, args := r.query(spec)

	err := r.requester.ExecOneWithScan(ctx, scanOneFunc(&dst), query, args...)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToSession(&dst), nil
}

func (r *PGReadRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*session.Session, error) {
	var dst = make([]*model.Sessions, 0)
	query, args := r.query(spec)

	err := r.requester.ExecManyWithScan(ctx, scanManyFunc(&dst), query, args)
	err = r.translateError(err)

	if err != nil && errx.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return ToSessions(dst), nil
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

	err = pg.WithEntity(err, session.ErrorSessionEntity)
	err = postgres.TranslateReadError(err)

	if pg.ErrorCode(err) == core.ErrorNotFoundCode {
		return nil
	}

	return err

}

/* __________________________________________________ */
