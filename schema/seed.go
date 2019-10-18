package schema

import (
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile("./seed.sql")
	if err != nil {
		return err
	}

	if _, err := tx.Exec(string(file)); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
