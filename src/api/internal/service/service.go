package service

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

type UserServiceInterface interface {
	Create(context.Context, domain.User) (domain.User, error)
	GetAll(context.Context) ([]domain.User, error)
}
