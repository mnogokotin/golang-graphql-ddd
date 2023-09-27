package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mnogokotin/golang-graphql-ddd/pkg/database/postgres"
	"io"
	"log"
	"os"

	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func main() {
	var err error

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseConnectionUri := postgres.GetConnectionUri()
	if len(databaseConnectionUri) == 0 {
		log.Fatalf("migrate: environment variables not declared")
	}

	db, err = sql.Open("postgres", databaseConnectionUri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("fixtures"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initiate new testfixture migrations_migrator: %v\n", err)
		os.Exit(1)
	}

	if err := fixtures.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load fixtures: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, "fixtures loaded\n")
}
