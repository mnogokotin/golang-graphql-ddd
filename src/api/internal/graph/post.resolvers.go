package graph

import (
	"context"
	"fmt"
	"github.com/mnogokotin/golang-graphql-ddd/internal/graph/model"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}
