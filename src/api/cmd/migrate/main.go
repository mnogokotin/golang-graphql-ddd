package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mnogokotin/golang-packages/database/postgres"
)

const (
	_defaultAttempts = 10
	_defaultTimeout  = time.Second
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseConnectionUri := postgres.GetConnectionUri()
	if len(databaseConnectionUri) == 0 {
		log.Fatalf("migrate: environment variables not declared")
	}

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseConnectionUri)
		if err == nil {
			break
		}

		log.Printf("migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate: no change")
		return
	}
	if err != nil {
		log.Fatalf("migrate: up error: %s", err)
	}

	log.Printf("migrate: up success")
}
