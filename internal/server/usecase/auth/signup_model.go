package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
)

/* __________________________________________________ */

type SignUpRequest struct {
	Payload SignUpRequestPayload
}

type SignUpRequestPayload struct {
	Credentials auth.CredentialsModel
}

/* __________________________________________________ */

type SignUpResponse struct {
	Payload SignUpResponsePayload
}

type SignUpResponsePayload struct {
	Session auth.SessionModel
	User    auth.UserModel
}

/* __________________________________________________ */
