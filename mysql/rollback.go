package mysql

import (
	"database/sql"
	"errors"
	"log"
)

func rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err == nil {
		return
	}

	if errors.Is(err, sql.ErrTxDone) {
		return
	}

	log.Println(err.Error())
}
