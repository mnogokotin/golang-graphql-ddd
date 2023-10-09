package service

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

type UserServiceInterface interface {
	GetAll(context.Context) ([]domain.User, error)
}
