package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
)

/* __________________________________________________ */

type SeekerAPI interface {
	GetByName(
		ctx context.Context, user string, name string,
	) (secret.EncodedSecret, error)
}

//goland:noinspection GoNameStartsWithPackageName
type SecretManagementAPI interface {
	CreateDecoded(
		ctx context.Context, user string, data secret.DecodedSecretData,
	) (secret.EncodedSecret, error)

	CreateEncoded(
		ctx context.Context, user string, data secret.EncodedSecretData,
	) (secret.EncodedSecret, error)
}

/* __________________________________________________ */
