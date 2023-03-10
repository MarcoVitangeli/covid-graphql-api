package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/MarcoVitangeli/covid-graphql-api/internal/cases"
)

// FindCases is the resolver for the findCases field.
func (r *queryResolver) FindCases(ctx context.Context, input *cases.CaseSearch) ([]*cases.Case, error) {
	return r.Resolver.srv.Find(ctx, input)
}

// FindByID is the resolver for the findById field.
func (r *queryResolver) FindByID(ctx context.Context, id *int) (*cases.Case, error) {
	c, err := r.Resolver.srv.Get(ctx, *id)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
