package repository

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-graphql-ddd/pkg/database/postgres"
)

type PostRepo struct {
	postgres *postgres.Postgres
}

func NewPostRepo(postgres *postgres.Postgres) *PostRepo {
	return &PostRepo{
		postgres: postgres,
	}
}

func (r *PostRepo) Store(ctx context.Context, u domain.Post) (domain.Post, error) {
	defer r.postgres.Close()

	result := r.postgres.Db.Create(&u)
	if result.Error != nil {
		return domain.Post{}, result.Error
	}

	return u, nil
}

func (r *PostRepo) GetAll(ctx context.Context) ([]domain.Post, error) {
	defer r.postgres.Close()

	var userModels []model.Post
	r.postgres.Db.Find(&userModels)

	var userDomains []domain.Post
	for _, m := range userModels {
		d := domain.Post{ID: m.ID, Text: m.Text}
		userDomains = append(userDomains, d)
	}

	return userDomains, nil
}
