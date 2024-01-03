package session

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/server/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gpl/search"
	"github.com/google/uuid"
)

/* __________________________________________________ */

type IDGenerator interface {
	Generate() session.SessionID
}

type IDGeneratorService struct{}

func ProvideIDGeneratorService() *IDGeneratorService {
	return &IDGeneratorService{}
}

func (s *IDGeneratorService) Generate() session.SessionID {
	return session.NewSessionID(uuid.NewString())
}

/* __________________________________________________ */

type Manager interface {
	Get(ctx context.Context, id common.UserID) (session.Session, error)
	Create(ctx context.Context, data session.SessionData) (session.Session, error)
}

type ManagerService struct {
	seq  IDGenerator
	repo Repository
}

func ProvideManagerService(
	seq IDGenerator,
	repo Repository,
) *ManagerService {
	return &ManagerService{
		seq:  seq,
		repo: repo,
	}
}

func (s *ManagerService) Get(
	ctx context.Context,
	id common.UserID,
) (session.Session, error) {
	ses, err := s.searchSession(ctx, id)
	if err != nil {
		return session.NewEmptySession(), err
	}
	return *ses, nil
}

func (s *ManagerService) Create(
	ctx context.Context,
	data session.SessionData,
) (session.Session, error) {

	id := s.seq.Generate()
	ses := *session.NewSession(id, data)

	result, err := s.repo.NonAutoSave(ctx, ses)
	if err != nil {
		return session.NewEmptySession(), err
	}

	return *result, err

}

func (s *ManagerService) searchSession(
	ctx context.Context,
	id common.UserID,
) (*session.Session, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(auth.UserIDCondition(id))
	return s.repo.FindOneBySpec(ctx, spec)
}

/* __________________________________________________ */
