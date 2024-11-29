package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	var migrationPath, connection string

	migrationPath, _ = os.LookupEnv("MIGRATION_PATH")
	connection, _ = os.LookupEnv("DB_CONNECTION")

	if migrationPath == "" || connection == "" {
		panic("MIGRATION_PATH and DB_CONNECTION env vars must be set")
	}

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	fsrc, err := (&file.File{}).Open(fmt.Sprintf("file://%s", migrationPath))
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no changes")
			return
		}
		panic(err)
	}
	fmt.Println("success")
}
