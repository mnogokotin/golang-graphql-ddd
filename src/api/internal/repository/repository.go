package repository

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

type UserRepoInterface interface {
	GetAll(context.Context) ([]domain.User, error)
}
