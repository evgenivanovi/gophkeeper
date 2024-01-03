package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
)

/* __________________________________________________ */

type SignUpRequestPayload struct {
	Credentials auth.CredentialsModel
}

type SignUpRequest struct {
	Payload SignUpRequestPayload
}

/* __________________________________________________ */

type SignUpResponsePayload struct {
	Session auth.SessionModel
	User    auth.UserModel
}

type SignUpResponse struct {
	Payload SignUpResponsePayload
}

/* __________________________________________________ */
