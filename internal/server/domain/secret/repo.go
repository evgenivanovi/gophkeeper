package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/search"
)

/* __________________________________________________ */

type KeyRepository interface {
	GetByID(
		ctx context.Context, id secret.SecretID,
	) (*secret.OwnedEncodedSecret, error)
	FindByID(
		ctx context.Context, id secret.SecretID,
	) (*secret.OwnedEncodedSecret, error)
	FindByIDs(
		ctx context.Context, ids []secret.SecretID,
	) ([]*secret.OwnedEncodedSecret, error)
}

type SpecificationRepository interface {
	FindOneBySpec(
		ctx context.Context, spec search.Specification,
	) (*secret.OwnedEncodedSecret, error)
	FindManyBySpec(
		ctx context.Context, spec search.Specification,
	) ([]*secret.OwnedEncodedSecret, error)
}

type ReadRepository interface {
	KeyRepository
	SpecificationRepository
}

type SaveRepository interface {
	AutoSave(
		ctx context.Context, data secret.OwnedEncodedSecretData, metadata core.Metadata,
	) (*secret.OwnedEncodedSecret, error)
}

type WriteRepository interface {
	SaveRepository
}

type Repository interface {
	ReadRepository
	WriteRepository
}

type RepositoryService struct {
	readRepository  ReadRepository
	writeRepository WriteRepository
}

func ProvideRepositoryService(
	readRepository ReadRepository,
	writeRepository WriteRepository,
) *RepositoryService {
	return &RepositoryService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

func (r *RepositoryService) GetByID(
	ctx context.Context, id secret.SecretID,
) (*secret.OwnedEncodedSecret, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *RepositoryService) FindByID(
	ctx context.Context, id secret.SecretID,
) (*secret.OwnedEncodedSecret, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *RepositoryService) FindByIDs(
	ctx context.Context, ids []secret.SecretID,
) ([]*secret.OwnedEncodedSecret, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *RepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*secret.OwnedEncodedSecret, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *RepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*secret.OwnedEncodedSecret, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *RepositoryService) AutoSave(
	ctx context.Context, data secret.OwnedEncodedSecretData, metadata core.Metadata,
) (*secret.OwnedEncodedSecret, error) {
	return r.writeRepository.AutoSave(ctx, data, metadata)
}

/* __________________________________________________ */
