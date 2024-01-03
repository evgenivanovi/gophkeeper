package secret

import (
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/model"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	timex "github.com/evgenivanovi/gophkeeper/pkg/time"
)

/* __________________________________________________ */

func ToSecret(record *model.Secrets) *secret.OwnedEncodedSecret {

	id := secret.NewSecretID(
		record.ID,
	)

	data := secret.NewOwnedEncodedSecretData(
		common.NewUserID(record.UserID),
		record.Name,
		ToSecretType(record.TypeID),
		record.Content,
	)

	metadata := core.NewMetadata(
		record.CreatedAt.UTC(),
		timex.UTC(record.UpdatedAt),
		timex.UTC(record.DeletedAt),
	)

	return secret.NewOwnedEncodedSecret(id, *data, *metadata)

}

func FromSecretData(data secret.OwnedEncodedSecretData, metadata core.Metadata) model.Secrets {
	return model.Secrets{
		// DATA
		UserID: data.UserID.ID(),
		TypeID: FromSecretType(data.Secret),

		Name:    data.Name,
		Content: data.Content,

		// METADATA
		CreatedAt: metadata.CreatedAt,
		UpdatedAt: metadata.UpdatedAt,
		DeletedAt: metadata.DeletedAt,
	}
}

func FromSecret(entity secret.OwnedEncodedSecret) model.Secrets {
	return model.Secrets{
		// ID
		ID: entity.Identity().ID(),

		// DATA
		UserID: entity.Data().UserID.ID(),
		TypeID: FromSecretType(entity.Data().Secret),

		Name:    entity.Data().Name,
		Content: entity.Data().Content,

		// METADATA
		CreatedAt: entity.Metadata().CreatedAt,
		UpdatedAt: entity.Metadata().UpdatedAt,
		DeletedAt: entity.Metadata().DeletedAt,
	}
}

func ToSecrets(records []*model.Secrets) []*secret.OwnedEncodedSecret {
	entities := make([]*secret.OwnedEncodedSecret, 0)
	for _, record := range records {
		entities = append(entities, ToSecret(record))
	}
	return entities
}

func FromSecrets(entities []secret.OwnedEncodedSecret) []model.Secrets {
	records := make([]model.Secrets, 0)
	for _, entity := range entities {
		records = append(records, FromSecret(entity))
	}
	return records
}

/* __________________________________________________ */

func ToSecretType(id int64) secret.SecretType {

	if id == 1 {
		return secret.Text
	} else if id == 2 {
		return secret.Binary
	} else if id == 3 {
		return secret.Credentials
	} else if id == 4 {
		return secret.Card
	}

	panic("invalid secret type")

}

func FromSecretType(sec secret.SecretType) int64 {

	if sec == secret.Text {
		return 1
	} else if sec == secret.Binary {
		return 2
	} else if sec == secret.Credentials {
		return 3
	} else if sec == secret.Card {
		return 4
	}

	panic("invalid secret type")

}

/* __________________________________________________ */
