package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mnogokotin/golang-graphql-ddd/graph/model"
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

func RunMigrations(db *gorm.DB) {
	if !db.HasTable(&model.Question{}) {
		db.CreateTable(&model.Question{})
	}
	if !db.HasTable(&model.Choice{}) {
		db.CreateTable(&model.Choice{})
		db.Model(&model.Choice{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	}
}
