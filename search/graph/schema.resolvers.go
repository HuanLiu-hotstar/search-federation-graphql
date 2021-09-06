package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/HuanLiu-hotstar/search-federation-graphql/search/graph/data"
	"github.com/HuanLiu-hotstar/search-federation-graphql/search/graph/generated"
	"github.com/HuanLiu-hotstar/search-federation-graphql/search/graph/model"
)

func (r *queryResolver) Search(ctx context.Context, w string) ([]*model.Content, error) {
	return data.FetchSearchData(w)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
