package database_test

import (
	"testing"

	"github.com/ochom/event-utils/database"
	"github.com/ochom/event-utils/utils"
)

func TestDatabase(t *testing.T) {
	t.Skip("TODO")
}

func initDB(t *testing.T) database.Repository {
	t.Helper()

	testEnv := utils.GetEnvOrDefault("TEST_ENV", "git")

	if testEnv == "local" {
		r, err := database.New(database.Postgres)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	r, err := database.New(database.SQLite)
	if err != nil {
		t.Fatal(err)
	}

	return r
}
