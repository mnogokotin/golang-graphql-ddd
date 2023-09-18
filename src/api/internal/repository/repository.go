package service

import (
	"context"

	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	GetAll(context.Context) ([]domain.User, error)
}
