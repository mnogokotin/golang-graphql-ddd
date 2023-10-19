package service

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

type PostServiceInterface interface {
	Create(context.Context, domain.Post) (domain.Post, error)
	GetAll(context.Context) ([]domain.Post, error)
}

type UserServiceInterface interface {
	Create(context.Context, domain.User) (domain.User, error)
	GetAll(context.Context) ([]domain.User, error)
}
