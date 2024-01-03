package session

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/token"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SessionID struct {
	id string
}

func NewSessionID(id string) SessionID {
	return SessionID{
		id: id,
	}
}

func (s SessionID) ID() string {
	return s.id
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SessionData struct {
	UserID common.UserID
	Token  token.Token
}

func NewSessionData(userID common.UserID, token token.Token) *SessionData {
	return &SessionData{
		UserID: userID,
		Token:  token,
	}
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type Session struct {
	id   SessionID
	data SessionData
}

func NewSession(id SessionID, data SessionData) *Session {
	return &Session{
		id:   id,
		data: data,
	}
}

func NewEmptySession() Session {
	return Session{}
}

func (e *Session) Identity() SessionID {
	return e.id
}

func (e *Session) Data() SessionData {
	return e.data
}

func ToSessionPointers(entities []Session) []*Session {
	result := make([]*Session, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToSessionValues(entities []*Session) []Session {
	result := make([]Session, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

/* __________________________________________________ */
