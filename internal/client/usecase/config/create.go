package config

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/fs"
)

/* __________________________________________________ */

type CreateConfigUsecase interface {
	Execute(ctx context.Context) error
}

type CreateConfigUsecaseService struct {
	fileManager fs.Manager
}

func ProvideCreateConfigUsecaseService(
	fileManager fs.Manager,
) *CreateConfigUsecaseService {
	return &CreateConfigUsecaseService{
		fileManager: fileManager,
	}
}

func (uc *CreateConfigUsecaseService) Execute(ctx context.Context) error {

	path := common.MustOptionsFromCtx(ctx).Config

	settings := fs.NewCreationSettings(
		path,
		fs.WithForceCreation(false),
	)

	err := uc.fileManager.Create(ctx, settings)
	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
