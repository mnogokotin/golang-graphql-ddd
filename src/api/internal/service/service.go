package service

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type User interface {
	GetAll(context.Context) ([]domain.User, error)
}
