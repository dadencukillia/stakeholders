// Shared logic for PostgreSQL connection, migrations etc
package db

import (
	"context"
	"embed"
	"fmt"

	"github.com/dadencukillia/stakeholders/shared/config"
	"github.com/dadencukillia/stakeholders/shared/db/sqlc"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Database struct {
	pool *pgxpool.Pool
	repo *sqlc.Queries
}

func ConnectConfigDatabase(ctx context.Context, cfg *config.ServiceConfig) (Database, error) {
	conURL := GetDBConnectionURL(
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		fmt.Sprint(cfg.Database.Port),
		cfg.Database.Name,
	)

	return ConnectDB(ctx, conURL)
}

func GetDBConnectionURL(user, password, host, port, database string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
}

func ConnectDB(ctx context.Context, connectionURL string) (Database, error) {
	conn, err := pgxpool.New(ctx, connectionURL)
	if err != nil {
		return Database{}, err
	}

	return Database{
		pool: conn,
		repo: sqlc.New(conn),
	}, nil
}

func (a Database) Close() {
	a.pool.Close()
}

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

func (a Database) GetRepo() *sqlc.Queries {
	return a.repo
}

func (a Database) Transaction(ctx context.Context, callback func(txRepo *sqlc.Queries) error) error {
	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("couldn't acquire a tx connection: %w", err)
	}

	defer tx.Rollback(context.Background())
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(context.Background())
			panic(p)
		}
	}()

	err = callback(a.repo.WithTx(tx))
	if err != nil {
		return fmt.Errorf("error from the transaction callback: %w", err)
	}

	if commitErr := tx.Commit(ctx); commitErr != nil {
		return fmt.Errorf("couldn't commit tx changes: %w", commitErr)
	}

	return nil
}
