package service

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
)

type PostService struct {
	r repository.PostRepoInterface
}

func NewPostService(r repository.PostRepoInterface) *PostService {
	return &PostService{
		r: r,
	}
}

func (s *PostService) Create(ctx context.Context, u domain.Post) (domain.Post, error) {
	postDomain, err := s.r.Store(ctx, u)
	if err != nil {
		return domain.Post{}, err
	}

	return postDomain, nil
}

func (s *PostService) GetAll(ctx context.Context) ([]domain.Post, error) {
	posts, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
