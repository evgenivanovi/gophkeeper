package secret

import (
	"github.com/evgenivanovi/gophkeeper/internal/server/postgres/public/model"
	"github.com/jackc/pgx/v5"
)

/* __________________________________________________ */

func scanOne(row pgx.Row, record *model.Secrets) error {
	return row.Scan(
		// ID
		&record.ID,
		// DATA
		&record.UserID,
		&record.TypeID,
		&record.Name,
		&record.Content,
		// DATA
		&record.CreatedAt,
		&record.UpdatedAt,
		&record.DeletedAt,
	)
}

func scanOneFunc(record *model.Secrets) func(row pgx.Row) error {
	return func(row pgx.Row) error {
		return scanOne(row, record)
	}
}

func scanMany(rows pgx.Rows, records *[]*model.Secrets) error {
	defer rows.Close()

	for rows.Next() {
		var record model.Secrets
		if err := scanOne(rows, &record); err != nil {
			return err
		}
		*records = append(*records, &record)
	}

	return nil
}

func scanManyFunc(records *[]*model.Secrets) func(rows pgx.Rows) error {
	return func(rows pgx.Rows) error {
		return scanMany(rows, records)
	}
}

/* __________________________________________________ */
