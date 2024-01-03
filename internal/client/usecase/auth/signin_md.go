package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
)

/* __________________________________________________ */

type SignInRequestPayload struct {
	Credentials auth.CredentialsModel
}

type SignInRequest struct {
	Payload SignInRequestPayload
}

/* __________________________________________________ */

type SignInResponsePayload struct {
	Session auth.SessionModel
	User    auth.UserModel
}

type SignInResponse struct {
	Payload SignInResponsePayload
}

/* __________________________________________________ */
