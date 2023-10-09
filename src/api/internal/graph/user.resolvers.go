package graph

import (
	"context"
	"fmt"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
	"github.com/mnogokotin/golang-graphql-ddd/internal/repository"
	"github.com/mnogokotin/golang-graphql-ddd/internal/service"
	"github.com/mnogokotin/golang-graphql-ddd/pkg/database/postgres"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, a int8) ([]*model.User, error) {
	pg, err := postgres.New(postgres.GetConnectionUri())
	if err != nil {
		return nil, err
	}

	userService := service.New(
		repository.New(pg),
	)

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
