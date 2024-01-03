package session

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gpl/search"
)

/* __________________________________________________ */

type KeyRepository interface {
	GetByID(ctx context.Context, id session.SessionID) (*session.Session, error)
	FindByID(ctx context.Context, id session.SessionID) (*session.Session, error)
	FindByIDs(ctx context.Context, ids []session.SessionID) ([]*session.Session, error)
}

type SpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*session.Session, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*session.Session, error)
}

type ReadRepository interface {
	KeyRepository
	SpecificationRepository
}

type SaveRepository interface {
	NonAutoSave(ctx context.Context, data session.Session) (*session.Session, error)
	NonAutoSaveAll(ctx context.Context, data []session.Session) error
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
	ctx context.Context, id session.SessionID,
) (*session.Session, error) {
	return r.readRepository.GetByID(ctx, id)
}

func (r *RepositoryService) FindByID(
	ctx context.Context, id session.SessionID,
) (*session.Session, error) {
	return r.readRepository.FindByID(ctx, id)
}

func (r *RepositoryService) FindByIDs(
	ctx context.Context, ids []session.SessionID,
) ([]*session.Session, error) {
	return r.readRepository.FindByIDs(ctx, ids)
}

func (r *RepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*session.Session, error) {
	return r.readRepository.FindOneBySpec(ctx, spec)
}

func (r *RepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*session.Session, error) {
	return r.readRepository.FindManyBySpec(ctx, spec)
}

func (r *RepositoryService) NonAutoSave(
	ctx context.Context, data session.Session,
) (*session.Session, error) {
	return r.writeRepository.NonAutoSave(ctx, data)
}

func (r *RepositoryService) NonAutoSaveAll(
	ctx context.Context, data []session.Session,
) error {
	return r.writeRepository.NonAutoSaveAll(ctx, data)
}

/* __________________________________________________ */
