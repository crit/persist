package mysql

import (
	"net/http"

	"github.com/crit/errors"
	"github.com/crit/persist"
)

func Delete(model persist.Deleter) error {
	tx, err := pool.Begin()
	if err != nil {
		return err
	}

	defer rollback(tx)

	res, err := tx.Exec(model.Query(), model.Values()...)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num > 1 {
		return errors.New(http.StatusBadRequest, "attempted to delete more than 1 row: %d", num)
	}

	return tx.Commit()
}
