package mysql

import "github.com/crit/persist"

func Find(model persist.Finder) error {
	tx, err := pool.Begin()
	if err != nil {
		return err
	}

	defer rollback(tx)

	row := tx.QueryRow(model.Query(), model.Values()...)

	err = model.Use(row)
	if err != nil {
		return err
	}

	return tx.Commit()
}
