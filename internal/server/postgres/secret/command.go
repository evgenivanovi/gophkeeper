package secret

import (
	"fmt"

	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/model"
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/table"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

/* __________________________________________________ */

func insertOneStatement(
	record model.Secrets,
) (string, []interface{}) {

	stmt := table.Secrets.
		INSERT(
			// DATA
			table.Secrets.UserID,
			table.Secrets.TypeID,
			table.Secrets.Name,
			table.Secrets.Content,
		).
		VALUES(
			// DATA
			record.UserID,
			record.TypeID,
			record.Name,
			record.Content,
		).
		RETURNING(
			table.Secrets.ID,
		)

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

func insertAllStatement(
	records []model.Secrets,
) (string, []interface{}) {

	stmt := table.Secrets.INSERT(
		// DATA
		table.Secrets.UserID,
		table.Secrets.TypeID,
		table.Secrets.Name,
		table.Secrets.Content,
	)

	for _, record := range records {
		stmt = stmt.VALUES(
			// DATA
			record.UserID,
			record.TypeID,
			record.Name,
			record.Content,
		)
	}

	slogx.Log().Debug(fmt.Sprintf("Calculated SQL query: `%s`", stmt.DebugSql()))
	return stmt.Sql()

}

/* __________________________________________________ */
