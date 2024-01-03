package session

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/session"
	"github.com/evgenivanovi/gpl/search"
	"github.com/evgenivanovi/gpl/stdx"
)

/* __________________________________________________ */

const IDSearchKey search.Key = "id"

func IdentityCondition(id session.SessionID) search.SearchCondition {
	return *search.NewEquality(
		IDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func IdentitiesCondition(ids []session.SessionID) search.SearchCondition {

	raws := make([]string, 0)
	for _, id := range ids {
		raws = append(raws, id.ID())
	}

	return *search.NewContainsAny(
		IDSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */
