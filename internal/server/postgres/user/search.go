package user

import (
	"fmt"

	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/table"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gpl/search"
	slices "github.com/evgenivanovi/gpl/std/slice"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	pgjet "github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func idExpression(
	id common.UserID,
) pgjet.BoolExpression {
	return table.Users.ID.EQ(pgjet.Int(id.ID()))
}

func idsExpression(
	ids []common.UserID,
) pgjet.BoolExpression {
	exp := make([]pgjet.Expression, 0)
	for _, id := range ids {
		exp = append(exp, pgjet.Int(id.ID()))
	}
	return table.Users.ID.IN(exp...)
}

func usernameExpression(
	username string,
) pgjet.BoolExpression {
	return table.Users.Username.EQ(pgjet.String(username))
}

func usernamesExpression(
	usernames []string,
) pgjet.BoolExpression {
	exp := make([]pgjet.Expression, 0)
	for _, username := range usernames {
		exp = append(exp, pgjet.String(username))
	}
	return table.Users.Username.IN(exp...)
}

/* __________________________________________________ */

func buildQuery(
	searchExp pgjet.BoolExpression,
	orderExp []pgjet.OrderByClause,
	lock pgjet.RowLock,
	sliceCondition search.SliceCondition,
) (string, []interface{}) {

	stmt := pgjet.
		SELECT(
			table.Users.AllColumns,
		).
		FROM(
			table.Users,
		)

	if sliceCondition.Chunked() {
		chunk := sliceCondition.Chunk()
		if limit, ok := chunk.Limit(); ok {
			stmt.LIMIT(limit)
		}
		if offset, ok := chunk.Offset(); ok {
			stmt.OFFSET(offset)
		}
	}

	if searchExp != nil {
		stmt = stmt.WHERE(searchExp)
	}

	if slices.IsNotEmpty(orderExp) {
		stmt = stmt.ORDER_BY(orderExp...)
	}

	if lock != nil {
		stmt = stmt.FOR(lock)
	}

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
