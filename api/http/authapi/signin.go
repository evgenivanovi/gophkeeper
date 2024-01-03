package authapi

import (
	"github.com/evgenivanovi/gophkeeper/api/http"
	"github.com/evgenivanovi/gophkeeper/api/http/auth"
)

/* __________________________________________________ */

var SigninEndpoint http.Endpoint = "/api/auth/signin"

/* __________________________________________________ */

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* __________________________________________________ */

type SigninResponsePayload struct {
	Session auth.SessionModel `json:"session"`
	User    auth.UserModel    `json:"user"`
}

type SigninResponse struct {
	Payload SigninResponsePayload `json:"payload"`
}

/* __________________________________________________ */
