// Shared logic for PostgreSQL connection, migrations etc
package db

import (
	"context"
	"fmt"

	"github.com/dadencukillia/stakeholders/shared/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
	repo *sqlc.Queries
}

func GetDBConnectionURL(username, password, host, port, database string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
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
