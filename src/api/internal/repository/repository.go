package repository

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

type UserRepoInterface interface {
	Store(context.Context, domain.User) (domain.User, error)
	GetAll(context.Context) ([]domain.User, error)
}
