package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func migrateDb(db *sql.DB) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("failed to create database migration driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"ql",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}
	version, dirty, _ := m.Version()
	log.Printf("migration database schema current version: %v, dirty: %v.\n", version, dirty)
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
