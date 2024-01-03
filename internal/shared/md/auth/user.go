package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
)

/* __________________________________________________ */

type UserModel struct {
	ID       int64
	Data     UserDataModel
	Metadata common.MetadataModel
}

type UserDataModel struct {
	Username string
}

/* __________________________________________________ */
