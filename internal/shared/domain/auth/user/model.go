package user

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type User struct {
	id       common.UserID
	data     UserData
	metadata core.Metadata
}

func NewUser(id common.UserID, data UserData, metadata core.Metadata) *User {
	return &User{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyUser() User {
	return User{}
}

func (e *User) Identity() common.UserID {
	return e.id
}

func (e *User) Data() UserData {
	return e.data
}

func (e *User) Metadata() core.Metadata {
	return e.metadata
}

func (e *User) ToAuthUser(data AuthUserData) *AuthUser {
	return NewAuthUser(e.id, data, e.metadata)
}

func ToUserPointers(entities []User) []*User {
	result := make([]*User, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToUserValues(entities []*User) []User {
	result := make([]User, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type UserData struct {
	Credentials auth.Credentials
}

func NewUserData(credentials auth.Credentials) *UserData {
	return &UserData{
		Credentials: credentials,
	}
}

/* __________________________________________________ */

type AuthUser struct {
	id       common.UserID
	data     AuthUserData
	metadata core.Metadata
}

func NewAuthUser(
	id common.UserID, data AuthUserData, metadata core.Metadata,
) *AuthUser {
	return &AuthUser{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyAuthUser() AuthUser {
	return AuthUser{}
}

func (e *AuthUser) Identity() common.UserID {
	return e.id
}

func (e *AuthUser) Data() AuthUserData {
	return e.data
}

func (e *AuthUser) Metadata() core.Metadata {
	return e.metadata
}

func (e *AuthUser) ToUser() *User {
	userData := *NewUserData(e.data.Credentials)
	return NewUser(e.id, userData, e.metadata)
}

/* __________________________________________________ */

type AuthUserData struct {
	SessionID   session.SessionID
	Tokens      token.Tokens
	Credentials auth.Credentials
}

func NewAuthUserData(
	sessionID session.SessionID, tokens token.Tokens, credentials auth.Credentials,
) *AuthUserData {
	return &AuthUserData{
		SessionID:   sessionID,
		Tokens:      tokens,
		Credentials: credentials,
	}
}

/* __________________________________________________ */
