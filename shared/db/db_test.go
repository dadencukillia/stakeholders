package db_test

import (
	"context"
	"errors"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/dadencukillia/stakeholders/shared/db"
	"github.com/dadencukillia/stakeholders/shared/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// DB connection

var testDB db.Database

func TestMain(m *testing.M) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start postgres container: %s", err)
	}

	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Printf("failed to terminate pgContainer: %s", err)
		}
	}()

	conURL, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}

	testDB, err = db.ConnectDB(ctx, conURL)
	if err != nil {
		log.Fatalf("failed to connect to test db: %s", err)
	}
	defer testDB.Close()

	if err := testDB.Migrate(); err != nil {
		log.Fatalf("failed to run migrations: %s", err)
	}

	os.Exit(m.Run())
}

// Tests

func TestTransactionInterrupt(t *testing.T) {
	if err := testDB.Migrate(); err != nil {
		t.Fatal(err)
	}

	var userID uuid.UUID
	var err error
	expectedError := errors.New("Ooops")

	err = testDB.Transaction(t.Context(), func(txRepo *sqlc.Queries) error {
		userID, err = txRepo.CreateUser(t.Context(), sqlc.CreateUserParams{
			UserName: pgtype.Text{
				String: "jackjack",
				Valid: true,
			},
			FullName: pgtype.Text{
				String: "Jack Jackson",
				Valid: true,
			},
		})

		if err != nil {
			t.Fatal(err)
		}

		return expectedError
	})
	if !errors.Is(err, expectedError) {
		t.Fatal(err)
	}

	_, err = testDB.GetRepo().GetUserById(t.Context(), userID)
	if err == nil {
		t.Fatal(errors.New("error expected"))
	}
}

func TestUserManipulations(t *testing.T) {
	repo := testDB.GetRepo()
	userID, err := repo.CreateUser(t.Context(), sqlc.CreateUserParams{
		UserName: pgtype.Text{
			String: "jackjack",
			Valid: true,
		},
		FullName: pgtype.Text{
			String: "Jack Jackson",
			Valid: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	userData, err := repo.GetUserById(t.Context(), userID)
	if err != nil {
		t.Fatal(err)
	}

	if !(reflect.DeepEqual(userData.UserName.String, "jackjack") && reflect.DeepEqual(userData.FullName.String, "Jack Jackson")) {
		t.Fatal(errors.New("invalid data in database"))
	}

	count, err := repo.DeleteUserById(t.Context(), userID)
	if err != nil {
		t.Fatal(err)
	}

	if count != 1 {
		t.Fatal(errors.New("invalid deletion count"))
	}

	userData, err = repo.GetUserById(t.Context(), userID)
	if err == nil {
		t.Fatal(errors.New("error expected"))
	}
}
