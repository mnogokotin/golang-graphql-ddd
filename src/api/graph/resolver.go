package graph

import "github.com/mnogokotin/golang-graphql-ddd/graph/model"

type Resolver struct {
	questions []*model.Question
	choices   []*model.Choice
}
