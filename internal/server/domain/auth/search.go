package auth

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gpl/search"
	"github.com/evgenivanovi/gpl/stdx"
)

/* __________________________________________________ */

const UserIDSearchKey search.Key = "user_id"

func UserIDCondition(id common.UserID) search.SearchCondition {
	return *search.NewEquality(
		UserIDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func UserIDsCondition(ids []common.UserID) search.SearchCondition {

	raws := make([]int64, 0)
	for _, id := range ids {
		raws = append(raws, id.ID())
	}

	return *search.NewContainsAny(
		UserIDSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */

const UsernameSearchKey search.Key = "username"

func UsernameCondition(username string) search.SearchCondition {
	return *search.NewEquality(
		UsernameSearchKey,
		stdx.NewValue(username),
	)
}

func UsernamesCondition(usernames []string) search.SearchCondition {

	raws := make([]string, 0)
	raws = append(raws, usernames...)

	return *search.NewContainsAny(
		UsernameSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */
