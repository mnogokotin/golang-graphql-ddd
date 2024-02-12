package service

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
)

type UserService struct {
	r repository.UserRepoInterface
}

func NewUserService(r repository.UserRepoInterface) *UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) Create(ctx context.Context, u domain.User) (domain.User, error) {
	userDomain, err := s.r.Store(ctx, u)
	if err != nil {
		return domain.User{}, err
	}

	return userDomain, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
