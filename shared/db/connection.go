// File with a couple of functions for database connection establishing or closing

package db

import (
	"context"
	"fmt"

	"github.com/dadencukillia/stakeholders/shared/config"
	"github.com/dadencukillia/stakeholders/shared/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Formats database runtime data into PostgreSQL connection string
func GetDBConnectionURL(user, password, host, port, database string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
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
