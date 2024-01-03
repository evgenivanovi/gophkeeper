package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	"github.com/evgenivanovi/gophkeeper/pkg/time"
)

/* __________________________________________________ */

func ToUserObject(
	userModel auth.UserModel,
	sessionModel auth.SessionModel,
) config.UserObject {
	return config.UserObject{
		User:    userModel.Data.Username,
		Session: sessionModel.ID,
		Secrets: config.SecretsObject{
			Access: config.SecretObject{
				Data:       sessionModel.Tokens.AccessToken.Token,
				Expiration: time.NewInstant(sessionModel.Tokens.AccessToken.ExpiresAt.UTC()),
			},
			Refresh: config.SecretObject{
				Data:       sessionModel.Tokens.RefreshToken.Token,
				Expiration: time.NewInstant(sessionModel.Tokens.RefreshToken.ExpiresAt.UTC()),
			},
		},
	}
}

/* __________________________________________________ */
