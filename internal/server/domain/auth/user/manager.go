package user

import (
	"context"

	authdm "github.com/evgenivanovi/gophkeeper/internal/server/domain/auth"
	sessiondm "github.com/evgenivanovi/gophkeeper/internal/server/domain/auth/session"
	tokendm "github.com/evgenivanovi/gophkeeper/internal/server/domain/auth/token"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	errx "github.com/evgenivanovi/gpl/err"
	"github.com/evgenivanovi/gpl/search"
)

/* __________________________________________________ */

type AuthManagerService struct {
	trx      common.Transactor
	repo     AuthRepository
	password PasswordManager
	token    tokendm.Manager
	session  sessiondm.Manager
}

func ProvideAuthManagerService(
	trx common.Transactor,
	repo AuthRepository,
	password PasswordManager,
	token tokendm.Manager,
	session sessiondm.Manager,
) *AuthManagerService {
	return &AuthManagerService{
		trx:      trx,
		repo:     repo,
		password: password,
		token:    token,
		session:  session,
	}
}

func (u *AuthManagerService) Signin(
	ctx context.Context,
	credentials auth.Credentials,
) (user.AuthUser, error) {

	principal, err := u.searchUser(ctx, credentials)
	if err != nil {
		return user.NewEmptyAuthUser(), err
	}

	err = u.checkPassword(ctx, credentials, *principal)
	if err != nil {
		return user.NewEmptyAuthUser(), err
	}

	return u.signinUser(ctx, *principal)

}

func (u *AuthManagerService) signinUser(
	ctx context.Context,
	principal user.User,
) (user.AuthUser, error) {

	tokens := u.generateTokens(ctx, principal)

	ses, err := u.createSession(ctx, principal, tokens)
	if err != nil {
		return user.NewEmptyAuthUser(), err
	}

	authUserData := *user.NewAuthUserData(
		ses.Identity(),
		tokens,
		principal.Data().Credentials,
	)

	authUser := *user.NewAuthUser(
		principal.Identity(),
		authUserData,
		principal.Metadata(),
	)

	return authUser, nil

}

func (u *AuthManagerService) Signup(
	ctx context.Context,
	credentials auth.Credentials,
) (user.User, error) {
	principal, err := u.createUser(ctx, credentials)
	if err != nil {
		return user.NewEmptyUser(), err
	}
	return *principal, nil
}

func (u *AuthManagerService) SignupAndSignin(
	ctx context.Context,
	credentials auth.Credentials,
) (user.AuthUser, error) {
	principal, err := u.Signup(ctx, credentials)
	if err != nil {
		return user.NewEmptyAuthUser(), err
	}
	return u.signinUser(ctx, principal)
}

func (u *AuthManagerService) createUser(
	ctx context.Context, credentials auth.Credentials,
) (*user.User, error) {

	data := user.NewUserData(
		credentials.WithHash(u.passwordHasher()),
	)

	principal, err := u.repo.AutoSave(
		ctx, *data, *core.NewNowMetadata(),
	)

	if err != nil {
		return nil, err
	}

	return principal, err

}

func (u *AuthManagerService) searchUser(
	ctx context.Context, credentials auth.Credentials,
) (*user.User, error) {

	spec := search.
		NewSpecificationTemplate().
		WithSearch(authdm.UsernameCondition(credentials.Username()))

	principal, err := u.repo.FindOneBySpec(ctx, spec)

	if err != nil {
		return nil, err
	}

	if principal == nil {
		return nil, errx.NewErrorWithEntityCode(user.ErrorUserEntity, core.ErrorNotFoundCode)
	}

	return principal, nil

}

func (u *AuthManagerService) checkPassword(
	ctx context.Context, credentials auth.Credentials, principal user.User,
) error {

	equaled := u.password.CompareHashPasswordCtx(
		ctx, credentials.Password(), principal.Data().Credentials.Password(),
	)

	if !equaled {
		return errx.NewErrorWithEntityCode(user.ErrorUserEntity, core.ErrorUnauthenticatedCode)
	}

	return nil

}

func (u *AuthManagerService) passwordHasher() func(string) string {
	return func(password string) string {
		hash, _ := u.
			password.
			GenerateHashPasswordCtx(context.Background(), password)
		return hash
	}
}

func (u *AuthManagerService) generateTokens(
	ctx context.Context, principal user.User,
) token.Tokens {
	data := token.NewTokenData(principal.Identity())
	return u.token.Generate(ctx, *data)
}

func (u *AuthManagerService) createSession(
	ctx context.Context, principal user.User, token token.Tokens,
) (session.Session, error) {
	data := session.NewSessionData(principal.Identity(), token.RefreshToken)
	return u.session.Create(ctx, *data)
}

/* __________________________________________________ */
