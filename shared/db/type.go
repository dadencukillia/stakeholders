package db

import (
	"github.com/dadencukillia/stakeholders/shared/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
	repo *sqlc.Queries
}
