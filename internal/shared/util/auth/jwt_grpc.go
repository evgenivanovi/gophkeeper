package auth

import (
	"context"
	"errors"

	"github.com/evgenivanovi/gpl/stdx/jwtx/grpc"
	"github.com/golang-jwt/jwt/v5"
)

/* __________________________________________________ */

// MetadataAuthKey ...
const MetadataAuthKey = "auth"

/* __________________________________________________ */

// GRPCExtractorProvider ...
func GRPCExtractorProvider() grpc.Extractor {
	return grpc.MultiExtractor{
		grpc.MetadataExtractor(MetadataAuthKey),
	}
}

// GRPCVerifierProvider ...
func GRPCVerifierProvider() func(context.Context, *jwt.Token, string) (context.Context, error) {
	return func(ctx context.Context, tkn *jwt.Token, token string) (context.Context, error) {

		if tkn == nil || !tkn.Valid {
			return nil, errors.New("token is invalid")
		}

		var claims *AccessClaims
		if tokenClaims, ok := tkn.Claims.(*AccessClaims); ok {
			claims = tokenClaims
		}

		if claims == nil || claims.User == nil || claims.User.UserID == 0 {
			return nil, errors.New("token is invalid")
		}

		return WithCtx(ctx, claims.User), nil

	}
}

/* __________________________________________________ */
