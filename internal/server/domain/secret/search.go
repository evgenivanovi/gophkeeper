package secret

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/search"
	"github.com/evgenivanovi/gpl/stdx"
)

/* __________________________________________________ */

const IDSearchKey search.Key = "id"

func IdentityCondition(id secret.SecretID) search.SearchCondition {
	return *search.NewEquality(
		IDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func IdentitiesCondition(ids []secret.SecretID) search.SearchCondition {

	raws := make([]int64, 0)
	for _, id := range ids {
		raws = append(raws, id.ID())
	}

	return *search.NewContainsAny(
		IDSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */

const NameSearchKey search.Key = "name"

func NameCondition(name string) search.SearchCondition {
	return *search.NewEquality(
		NameSearchKey,
		stdx.NewValue(name),
	)
}

func NamesCondition(names []string) search.SearchCondition {

	raws := make([]string, 0)
	for _, name := range names {
		raws = append(raws, name)
	}

	return *search.NewContainsAny(
		NameSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */
