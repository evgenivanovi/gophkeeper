package common

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
)

/* __________________________________________________ */

func FromMetadata(entity core.Metadata) MetadataModel {
	return MetadataModel{
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func ToMetadata(model MetadataModel) core.Metadata {
	return core.Metadata{
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}

/* __________________________________________________ */
