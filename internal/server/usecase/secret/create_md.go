package secret

import (
	secretmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type CreateDecodedSecretRequest struct {
	Payload CreateDecodedSecretRequestPayload
}

type CreateDecodedSecretRequestPayload struct {
	Secret secretmd.OwnedDecodedSecretDataModel
}

type CreateEncodedSecretRequest struct {
	Payload CreateEncodedSecretRequestPayload
}

type CreateEncodedSecretRequestPayload struct {
	Secret secretmd.OwnedEncodedSecretDataModel
}

/* __________________________________________________ */

type CreateDecodedSecretResponse struct {
	Payload CreateDecodedSecretResponsePayload
}

type CreateDecodedSecretResponsePayload struct {
	Secret secretmd.OwnedDecodedSecretModel
}

type CreateEncodedSecretResponse struct {
	Payload CreateEncodedSecretResponsePayload
}

type CreateEncodedSecretResponsePayload struct {
	Secret secretmd.OwnedEncodedSecretModel
}

/* __________________________________________________ */
