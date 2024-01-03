package auth

/* __________________________________________________ */

type SessionModel struct {
	ID     string      `json:"id"`
	Tokens TokensModel `json:"tokens"`
}

/* __________________________________________________ */
