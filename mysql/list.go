package mysql

import "github.com/crit/persist"

type Limiter struct {
	Limit  int `json:"limit"  validate:"gt=0"`
	Offset int `json:"offset" validate:"gte=0"`
}

type LimitMeta struct {
	Limiter
	Total int `json:"total"`
}

func List(model persist.Lister) error {
	tx, err := pool.Begin()
	if err != nil {
		return err
	}

	defer rollback(tx)

	rows, err := tx.Query(model.Query(), model.Values()...)
	if err != nil {
		return err
	}

	defer rows.Close()

	err = model.UseEach(rows)
	if err != nil {
		return err
	}

	var count int64
	err = tx.QueryRow("SELECT FOUND_ROWS() as Count").Scan(&count)
	if err != nil {
		return err
	}

	model.SetTotal(count)

	return tx.Commit()
}
