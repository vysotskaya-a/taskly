package pg

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// Init инициализирует Postgres db.
func Init(pgURL string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", pgURL)
	if err != nil {
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		os.Exit(1)
	}

	return db
}
