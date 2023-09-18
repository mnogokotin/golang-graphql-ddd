package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_SSLMODE"))
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getDatabaseUrl())
	return db, err
}
