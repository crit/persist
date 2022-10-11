package persist

import (
	"database/sql"
	"time"
)

type (
	HasQuery interface {
		Query() string
	}

	HasValues interface {
		Values() []any
	}

	NeedsInsertID interface {
		SetID(id int64)
	}

	NeedsRow interface {
		Use(row *sql.Row) error
	}

	NeedsRows interface {
		UseEach(rows *sql.Rows) error
		SetTotal(count int64)
	}

	NeedsAffectedCount interface {
		Affected(count int64) error
	}

	Deleter interface {
		HasQuery
		HasValues
	}

	Finder interface {
		HasQuery
		HasValues
		NeedsRow
	}

	Inserter interface {
		HasQuery
		HasValues
		NeedsInsertID
	}

	Joiner interface {
		HasQuery
		HasValues
	}

	Lister interface {
		HasQuery
		HasValues
		NeedsRows
	}

	Updater interface {
		HasQuery
		HasValues
		NeedsAffectedCount
	}

	Scanner func(dest ...any) error

	Tracking struct {
		CreatedAt time.Time `json:"created_at"`
		CreatedBy string    `json:"created_by"`
		UpdatedAt time.Time `json:"updated_at"`
		UpdatedBy string    `json:"updated_by"`
	}
)

const TrackingFields = `CreatedAt, CreatedBy, UpdatedAt, UpdatedBy`

func (Tracking) Fields() string {
	return TrackingFields
}
