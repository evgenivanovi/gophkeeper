package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
)

/* __________________________________________________ */

type SignInRequest struct {
	Payload SignInRequestPayload
}

type SignInRequestPayload struct {
	Credentials auth.CredentialsModel
}

/* __________________________________________________ */

type SignInResponse struct {
	Payload SignInResponsePayload
}

type SignInResponsePayload struct {
	Session auth.SessionModel
	User    auth.UserModel
}

/* __________________________________________________ */
