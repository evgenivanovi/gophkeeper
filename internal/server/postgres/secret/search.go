package secret

import (
	"fmt"

	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/table"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/search"
	slices "github.com/evgenivanovi/gpl/std/slice"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	pgjet "github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func idExpression(
	id secret.SecretID,
) pgjet.BoolExpression {
	return table.Secrets.ID.EQ(pgjet.Int(id.ID()))
}

func idsExpression(
	ids []secret.SecretID,
) pgjet.BoolExpression {
	exp := make([]pgjet.Expression, 0)
	for _, id := range ids {
		exp = append(exp, pgjet.Int(id.ID()))
	}
	return table.Secrets.ID.IN(exp...)
}

func userExpression(
	id common.UserID,
) pgjet.BoolExpression {
	return table.Secrets.UserID.EQ(pgjet.Int(id.ID()))
}

func usersExpression(
	ids []common.UserID,
) pgjet.BoolExpression {
	exp := make([]pgjet.Expression, 0)
	for _, id := range ids {
		exp = append(exp, pgjet.Int(id.ID()))
	}
	return table.Secrets.UserID.IN(exp...)
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
			table.Secrets.AllColumns,
		).
		FROM(
			table.Secrets,
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
