package fs

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/pkg/file"
)

/* __________________________________________________ */

type Manager interface {
	Create(ctx context.Context, settings CreationSettings) error
}

type ManagerService struct{}

func ProvideManagerService() *ManagerService {
	return &ManagerService{}
}

func (m *ManagerService) Create(
	ctx context.Context, settings CreationSettings,
) error {

	exists := file.Exists(settings.Path)

	if !exists {
		return file.Create(settings.Path)
	}

	if exists && settings.Options.Force {
		return file.CreateForce(settings.Path)
	}

	return nil

}

/* __________________________________________________ */
