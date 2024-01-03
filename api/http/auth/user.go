package auth

import (
	"github.com/evgenivanovi/gophkeeper/api/http/common"
)

/* __________________________________________________ */

type UserModel struct {
	ID       int64                `json:"id"`
	Metadata common.MetadataModel `json:"metadata"`
}

/* __________________________________________________ */
