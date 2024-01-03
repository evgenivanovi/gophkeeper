package auth

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type AuthAPI interface {
	Signin(
		ctx context.Context, credentials auth.Credentials,
	) (user.AuthUser, error)

	Signup(
		ctx context.Context, credentials auth.Credentials,
	) (user.User, error)

	SignupAndSignin(
		ctx context.Context, credentials auth.Credentials,
	) (user.AuthUser, error)
}

/* __________________________________________________ */
