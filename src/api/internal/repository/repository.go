package repository

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

type PostRepoInterface interface {
	Store(context.Context, domain.Post) (domain.Post, error)
	GetAll(context.Context) ([]domain.Post, error)
}

type UserRepoInterface interface {
	Store(context.Context, domain.User) (domain.User, error)
	GetAll(context.Context) ([]domain.User, error)
}
