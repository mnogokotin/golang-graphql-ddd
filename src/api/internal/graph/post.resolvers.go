package graph

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
	"github.com/mnogokotin/golang-graphql-ddd/internal/service"
	"github.com/mnogokotin/golang-graphql-ddd/pkg/database/postgres"
)

func createPostService() (*service.PostService, error) {
	pg, err := postgres.New(postgres.GetConnectionUri())
	if err != nil {
		return nil, err
	}

	return service.NewPostService(
		repository.NewPostRepo(pg),
	), nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	userService, err := createPostService()
	if err != nil {
		return nil, err
	}

	d := &domain.Post{Text: input.Text}
	d2, err := userService.Create(ctx, *d)
	if err != nil {
		return nil, err
	}

	return &model.Post{ID: d2.ID, Text: d2.Text}, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	userService, err := createPostService()
	if err != nil {
		return nil, err
	}

	userDomains, err := userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var userModels []*model.Post
	for _, d := range userDomains {
		m := &model.Post{ID: d.ID, Text: d.Text}
		userModels = append(userModels, m)
	}

	return userModels, nil
}
