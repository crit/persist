package mysql

import "database/sql"

var pool *sql.DB

func Register(instance *sql.DB) {
	if instance == nil {
		panic("provided sql.DB instance was nil")
	}

	if pool != nil {
		panic("attempted second registration of sql.DB instance")
	}

	pool = instance
}
