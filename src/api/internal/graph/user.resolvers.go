package graph

import (
	"context"
	"github.com/mnogokotin/golang-graphql-ddd/internal/domain"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
	"github.com/mnogokotin/golang-graphql-ddd/internal/service"
	"github.com/mnogokotin/golang-graphql-ddd/pkg/database/postgres"
)

func createUserService() (*service.UserService, error) {
	pg, err := postgres.New(postgres.GetConnectionUri())
	if err != nil {
		return nil, err
	}

	return service.NewUserService(
		repository.NewUserRepo(pg),
	), nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	userService, err := createUserService()
	if err != nil {
		return nil, err
	}

	d := &domain.User{Name: input.Name, Email: input.Email}
	d2, err := userService.Create(ctx, *d)
	if err != nil {
		return nil, err
	}

	return &model.User{ID: d2.ID, Name: d2.Name, Email: d2.Email}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userService, err := createUserService()
	if err != nil {
		return nil, err
	}

	userDomains, err := userService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var userModels []*model.User
	for _, d := range userDomains {
		m := &model.User{ID: d.ID, Name: d.Name, Email: d.Email}
		userModels = append(userModels, m)
	}

	return userModels, nil
}
