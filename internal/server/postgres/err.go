package postgres

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	errx "github.com/evgenivanovi/gpl/err"
	"github.com/evgenivanovi/gpl/pg"
)

/* __________________________________________________ */

func TranslateReadError(err error) error {

	if err == nil {
		return nil
	}

	code := pg.ErrorCode(err)
	entity := pg.ErrorEntity(err)

	if code == pg.ErrorEmptyCode {
		return errx.NewErrorWithEntityCode(
			entity, core.ErrorNotFoundCode,
		)
	}

	return errx.NewErrorWithEntityCodeMessage(
		entity, errx.ErrorInternalCode, err.Error(),
	)

}

func TranslateWriteError(err error) error {

	if err == nil {
		return nil
	}

	code := pg.ErrorCode(err)
	entity := pg.ErrorEntity(err)

	if code == pg.ErrorUniqueCode {
		return errx.NewErrorWithEntityCode(
			entity, core.ErrorExistsCode,
		)
	}

	return errx.NewErrorWithEntityCodeMessage(
		entity, errx.ErrorInternalCode, err.Error(),
	)

}

/* __________________________________________________ */
