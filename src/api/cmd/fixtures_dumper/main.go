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
	db *sql.DB
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

	dumper, err := testfixtures.NewDumper(
		testfixtures.DumpDatabase(db),
		testfixtures.DumpDialect("postgres"),
		testfixtures.DumpDirectory("fixtures"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initiate new testfixture dumper: %v\n", err)
		os.Exit(1)
	}
	if err := dumper.Dump(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to dump db: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, "db dumped\n")
}
