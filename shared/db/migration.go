package db

import (
	"embed"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (a Database) Migrate() error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(a.pool)
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return db.Close()
}
