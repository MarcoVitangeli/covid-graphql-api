package graph

import (
	"context"
	"github.com/MarcoVitangeli/covid-graphql-api/internal/cases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type service interface {
	Get(ctx context.Context, id int) (cases.Case, error)
	Find(ctx context.Context, query *cases.CaseSearch) ([]*cases.Case, error)
}
type Resolver struct {
	srv service
}

func NewResolver(s service) *Resolver {
	return &Resolver{
		srv: s,
	}
}
