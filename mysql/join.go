package mysql

import (
	"net/http"
	"strings"

	"github.com/crit/errors"
	"github.com/crit/persist"
)

func Join(model persist.Joiner) error {
	tx, err := pool.Begin()
	if err != nil {
		return err
	}

	defer rollback(tx)

	_, err = tx.Exec(model.Query(), model.Values()...)
	if err == nil {
		return tx.Commit()
	}

	if strings.Contains(err.Error(), "Error 1062") {
		return nil // we ignore duplication attempts to support idempotent calls
	}

	if strings.Contains(err.Error(), "Error 1452") {
		return errors.New(http.StatusBadRequest, "cannot join due to missing parent entity")
	}

	return err
}
