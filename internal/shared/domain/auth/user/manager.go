package user

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
)

/* __________________________________________________ */

type AuthManager interface {
	Signin(
		ctx context.Context,
		credentials auth.Credentials,
	) (AuthUser, error)

	Signup(
		ctx context.Context,
		credentials auth.Credentials,
	) (User, error)

	SignupAndSignin(
		ctx context.Context,
		credentials auth.Credentials,
	) (AuthUser, error)
}

/* __________________________________________________ */
