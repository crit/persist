package mysql

import "github.com/crit/persist"

func Insert(model persist.Inserter) error {
	tx, err := pool.Begin()
	if err != nil {
		return err
	}

	defer rollback(tx)

	res, err := tx.Exec(model.Query(), model.Values()...)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	model.SetID(id)

	return tx.Commit()
}
