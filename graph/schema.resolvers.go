package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"github.com/MarcoVitangeli/covid-graphql-api/graph/model"
)

// FindCases is the resolver for the findCases field.
func (r *queryResolver) FindCases(ctx context.Context, input *model.CaseSearch) ([]*model.Case, error) {
	return []*model.Case{&model.Case{
		ID:           "1234",
		Province:     "CABA",
		Gender:       "MALE",
		Neighborhood: "La Boca",
		Age:          20,
		Stage:        "CONFIRMADO",
		Dead:         "SI",
	}}, nil
}

// FindByID is the resolver for the findById field.
func (r *queryResolver) FindByID(ctx context.Context, id *int) (*model.Case, error) {
	return &model.Case{
		ID:           "1234",
		Province:     "CABA",
		Gender:       "MALE",
		Neighborhood: "La Boca",
		Age:          20,
		Stage:        "CONFIRMADO",
		Dead:         "SI",
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
