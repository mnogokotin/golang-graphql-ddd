package repository

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-graphql-ddd/pkg/database/postgres"
)

type UserRepo struct {
	postgres *postgres.Postgres
}

func New(postgres *postgres.Postgres) *UserRepo {
	return &UserRepo{
		postgres: postgres,
	}
}

func (r *UserRepo) GetAll(ctx context.Context) ([]domain.User, error) {
	defer r.postgres.Close()

	var userModels []model.User
	r.postgres.Db.Find(&userModels)

	var userDomains []domain.User
	for _, m := range userModels {
		d := domain.User{ID: m.ID, Name: m.Name, Email: m.Email}
		userDomains = append(userDomains, d)
	}

	return userDomains, nil
}
