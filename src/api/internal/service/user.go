package service

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
)

type UserService struct {
	repo repository.User
}

func New(r User) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
