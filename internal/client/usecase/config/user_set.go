package config

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
)

/* __________________________________________________ */

type SetUserUsecase interface {
	Execute(ctx context.Context, current string) error
}

type SetUserUsecaseService struct {
	configManager config.Manager
}

func ProvideSetUserUsecaseService(
	configManager config.Manager,
) *SetUserUsecaseService {
	return &SetUserUsecaseService{
		configManager: configManager,
	}
}

func (uc *SetUserUsecaseService) Execute(
	ctx context.Context,
	current string,
) error {
	path := common.MustOptionsFromCtx(ctx).Config
	return uc.configManager.SetCurrentUser(ctx, path, current)
}

/* __________________________________________________ */
