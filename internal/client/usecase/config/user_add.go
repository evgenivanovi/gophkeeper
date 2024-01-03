package config

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/fs"
)

/* __________________________________________________ */

type AddUserUsecase interface {
	Execute(ctx context.Context, user config.UserObject) error
}

type AddUserUsecaseService struct {
	fileManager   fs.Manager
	configManager config.Manager
}

func ProvideAddUserUsecaseService(
	fileManager fs.Manager,
	configManager config.Manager,
) *AddUserUsecaseService {
	return &AddUserUsecaseService{
		fileManager:   fileManager,
		configManager: configManager,
	}
}

func (uc *AddUserUsecaseService) Execute(
	ctx context.Context,
	user config.UserObject,
) error {
	path := common.MustOptionsFromCtx(ctx).Config
	return uc.configManager.AddUser(ctx, path, user)
}

/* __________________________________________________ */
