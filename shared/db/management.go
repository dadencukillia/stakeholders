// File with functions that return interfaces for database management

package db

import (
	"context"
	"fmt"

	"github.com/dadencukillia/stakeholders/shared/db/sqlc"
)

func (a Database) GetRepo() *sqlc.Queries {
	return a.repo
}

// Commits all databse operations in callback at once ensuring atomicity
// Or rollbacks all changes if error occured
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
