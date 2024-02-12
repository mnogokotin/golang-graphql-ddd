package repository

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-packages/database/postgres"
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

	var postModels []model.Post
	r.postgres.Db.Find(&postModels)
	var postDomains []domain.Post

	for _, m := range postModels {
		d := domain.Post{ID: m.ID, Text: m.Text, UserID: m.UserID}
		postDomains = append(postDomains, d)
	}

	return postDomains, nil
}
