package authapi

import (
	"github.com/evgenivanovi/gophkeeper/api/http"
	"github.com/evgenivanovi/gophkeeper/api/http/auth"
)

/* __________________________________________________ */

var SignupEndpoint http.Endpoint = "/api/auth/signup"

/* __________________________________________________ */

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* __________________________________________________ */

type SignupResponsePayload struct {
	Session auth.SessionModel `json:"session"`
	User    auth.UserModel    `json:"user"`
}

type SignupResponse struct {
	Payload SignupResponsePayload `json:"payload"`
}

/* __________________________________________________ */
