package graph

import (
	"context"
	"errors"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
	"github.com/mnogokotin/golang-graphql-ddd/internal/service"
	"github.com/mnogokotin/golang-packages/database/postgres"
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

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	userService, err := createUserService()
	if err != nil {
		return nil, err
	}

	userDomains, err := userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, ud := range userDomains {
		if ud.ID == obj.UserID {
			return &model.User{ID: ud.ID, Name: ud.Name, Email: ud.Email}, nil
		}
	}

	return &model.User{}, errors.New("post's user not found")
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	postService, err := createPostService()
	if err != nil {
		return nil, err
	}

	d := &domain.Post{Text: input.Text, UserID: input.UserID}
	d2, err := postService.Create(ctx, *d)
	if err != nil {
		return nil, err
	}

	return &model.Post{ID: d2.ID, Text: d2.Text}, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	postService, err := createPostService()
	if err != nil {
		return nil, err
	}

	postDomains, err := postService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var postModels []*model.Post
	for _, d := range postDomains {
		m := &model.Post{ID: d.ID, Text: d.Text, UserID: d.UserID}
		postModels = append(postModels, m)
	}

	return postModels, nil
}
