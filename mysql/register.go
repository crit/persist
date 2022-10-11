package mysql

import (
	"database/sql"
	"net/http"

	"github.com/crit/errors"
	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB

func Register(instance *sql.DB) error {
	if instance == nil {
		return errors.New(http.StatusInternalServerError, "provided sql.DB instance was nil")
	}

	if pool != nil {
		return errors.New(http.StatusInternalServerError, "attempted second registration of sql.DB instance")
	}

	pool = instance
	return nil
}
