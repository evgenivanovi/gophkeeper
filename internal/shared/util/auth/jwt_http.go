package auth

import (
	"net/http"

	"github.com/evgenivanovi/gpl/stdx/jwtx"
	jwthttp "github.com/evgenivanovi/gpl/stdx/jwtx/http"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

// CookieAuthKey ...
const CookieAuthKey = "auth"

/* __________________________________________________ */

// HTTPExtractorProvider ...
func HTTPExtractorProvider() request.Extractor {
	return request.MultiExtractor{
		jwthttp.CookieExtractor(
			CookieAuthKey,
		),
		request.HeaderExtractor{
			headers.AuthorizationKey.String(),
		},
	}
}

// HTTPVerifierProvider ...
func HTTPVerifierProvider() func(http.ResponseWriter, *http.Request, *jwt.Token, string) error {
	return func(writer http.ResponseWriter, request *http.Request, tkn *jwt.Token, token string) error {

		if tkn == nil || !tkn.Valid {
			return errors.New("token is invalid")
		}

		var claims *AccessClaims
		if tokenClaims, ok := tkn.Claims.(*AccessClaims); ok {
			claims = tokenClaims
		}

		if claims == nil || claims.User == nil || claims.User.UserID == 0 {
			return errors.New("token is invalid")
		}

		WithRequestCtx(request, claims.User)
		WriteTokenToResponseCookie(writer, tkn, token)

		return nil

	}
}

/* __________________________________________________ */

// WriteTokenToRequestContext ...
func WriteTokenToRequestContext(request *http.Request, tkn *jwt.Token, token string) {
	ctx := request.Context()
	if tkn != nil {
		ctx = jwtx.WithCtx(ctx, tkn)
	}
	if token != "" {
		ctx = jwtx.WithCtxAsString(ctx, token)
	}
	*request = *request.WithContext(ctx)
}

// WriteTokenToRequestCookie ...
func WriteTokenToRequestCookie(request *http.Request, tkn *jwt.Token, token string) {

	tokenExpirationTime, _ := tkn.Claims.GetExpirationTime()
	cookieExpirationTime := int(tokenExpirationTime.Unix())

	cookie := &http.Cookie{
		Name:   CookieAuthKey,
		Value:  token,
		MaxAge: cookieExpirationTime,
	}

	request.AddCookie(cookie)

}

// WriteTokenToResponseCookie ...
func WriteTokenToResponseCookie(writer http.ResponseWriter, tkn *jwt.Token, token string) {

	tokenExpirationTime, _ := tkn.Claims.GetExpirationTime()
	cookieExpirationTime := int(tokenExpirationTime.Unix())

	cookie := &http.Cookie{
		Name:   CookieAuthKey,
		Value:  token,
		MaxAge: cookieExpirationTime,
	}

	http.SetCookie(writer, cookie)

}

/* __________________________________________________ */
