package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/search"
)

/* __________________________________________________ */

type Seeker interface {
	GetByID(
		ctx context.Context, id secret.SecretID,
	) (secret.OwnedEncodedSecret, error)

	GetByName(
		ctx context.Context, user common.UserID, name string,
	) (secret.OwnedEncodedSecret, error)
}

type SeekerService struct {
	repo Repository
	enc  secret.SecretContentEncoder
	dec  secret.SecretContentDecoder
}

func ProvideSeekerService(
	repo Repository,
	enc secret.SecretContentEncoder,
	dec secret.SecretContentDecoder,
) *SeekerService {
	return &SeekerService{
		repo: repo,
		enc:  enc,
		dec:  dec,
	}
}

func (m *SeekerService) GetByID(
	ctx context.Context, id secret.SecretID,
) (secret.OwnedEncodedSecret, error) {
	sec, err := m.searchSessionByID(ctx, id)
	if err != nil {
		return secret.NewEmptyOwnedEncodedSecret(), err
	}
	return *sec, nil
}

func (m *SeekerService) GetByName(
	ctx context.Context, user common.UserID, name string,
) (secret.OwnedEncodedSecret, error) {
	sec, err := m.searchSessionByUserAndName(ctx, user, name)
	if err != nil {
		return secret.NewEmptyOwnedEncodedSecret(), err
	}
	return *sec, nil
}

func (m *SeekerService) searchSessionByID(
	ctx context.Context,
	id secret.SecretID,
) (*secret.OwnedEncodedSecret, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(IdentityCondition(id))
	return m.repo.FindOneBySpec(ctx, spec)
}

func (m *SeekerService) searchSessionByUserAndName(
	ctx context.Context,
	user common.UserID,
	name string,
) (*secret.OwnedEncodedSecret, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(auth.UserIDCondition(user)).
		WithSearch(NameCondition(name))
	return m.repo.FindOneBySpec(ctx, spec)
}

/* __________________________________________________ */
