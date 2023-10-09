package service

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
)

type UserService struct {
	r repository.UserRepoInterface
}

func New(r repository.UserRepoInterface) *UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
