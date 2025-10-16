package config

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB, embedMigrations embed.FS) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(db, "internal/migrations"); err != nil {
		return err
	}

	return nil
}
