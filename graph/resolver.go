package graph

import (
	"github.com/arthur1470/fc2-graphql/graph/generated"
	"github.com/arthur1470/fc2-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}

func (r *Resolver) Directive() generated.DirectiveResolver {
	//TODO implement me
	panic("implement me")
}

func (r *Resolver) Schema() generated.SchemaResolver {
	//TODO implement me
	panic("implement me")
}

func (r *Resolver) Type() generated.TypeResolver {
	//TODO implement me
	panic("implement me")
}
