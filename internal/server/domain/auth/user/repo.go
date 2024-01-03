package user

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gpl/search"
	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

type AuthKeyRepository interface {
	GetByID(ctx context.Context, id common.UserID) (*user.User, error)
	FindByID(ctx context.Context, id common.UserID) (*user.User, error)
	FindByIDs(ctx context.Context, ids []common.UserID) ([]*user.User, error)
}

type AuthSpecificationRepository interface {
	FindOneBySpec(ctx context.Context, spec search.Specification) (*user.User, error)
	FindManyBySpec(ctx context.Context, spec search.Specification) ([]*user.User, error)
}

type AuthReadRepository interface {
	AuthKeyRepository
	AuthSpecificationRepository
}

/* __________________________________________________ */

type AuthSaveRepository interface {
	AutoSave(ctx context.Context, data user.UserData, metadata core.Metadata) (*user.User, error)
	AutoSaveAll(ctx context.Context, datas []std.Pair[user.UserData, core.Metadata]) error
}

type AuthWriteRepository interface {
	AuthSaveRepository
}

/* __________________________________________________ */

type AuthRepository interface {
	AuthReadRepository
	AuthWriteRepository
}

type AuthRepositoryService struct {
	read  AuthReadRepository
	write AuthWriteRepository
}

func ProvideAuthRepositoryService(
	read AuthReadRepository,
	write AuthWriteRepository,
) *AuthRepositoryService {
	return &AuthRepositoryService{
		read:  read,
		write: write,
	}
}

func (r *AuthRepositoryService) GetByID(
	ctx context.Context, id common.UserID,
) (*user.User, error) {
	return r.read.GetByID(ctx, id)
}

func (r *AuthRepositoryService) FindByID(
	ctx context.Context, id common.UserID,
) (*user.User, error) {
	return r.read.FindByID(ctx, id)
}

func (r *AuthRepositoryService) FindByIDs(
	ctx context.Context, ids []common.UserID,
) ([]*user.User, error) {
	return r.read.FindByIDs(ctx, ids)
}

func (r *AuthRepositoryService) FindOneBySpec(
	ctx context.Context, spec search.Specification,
) (*user.User, error) {
	return r.read.FindOneBySpec(ctx, spec)
}

func (r *AuthRepositoryService) FindManyBySpec(
	ctx context.Context, spec search.Specification,
) ([]*user.User, error) {
	return r.read.FindManyBySpec(ctx, spec)
}

func (r *AuthRepositoryService) AutoSave(
	ctx context.Context, data user.UserData, metadata core.Metadata,
) (*user.User, error) {
	return r.write.AutoSave(ctx, data, metadata)
}

func (r *AuthRepositoryService) AutoSaveAll(
	ctx context.Context, datas []std.Pair[user.UserData, core.Metadata],
) error {
	return r.write.AutoSaveAll(ctx, datas)
}

/* __________________________________________________ */
