package mysql

import "github.com/crit/persist"

func Update(model persist.Updater) error {
	tx, err := pool.Begin()
	if err != nil {
		return err
	}

	defer rollback(tx)

	res, err := tx.Exec(model.Query(), model.Values()...)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	err = model.Affected(count)

	if err != nil {
		return err
	}

	return tx.Commit()
}
