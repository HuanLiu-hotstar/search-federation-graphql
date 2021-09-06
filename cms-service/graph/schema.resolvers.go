package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/HuanLiu-hotstar/search-federation-graphql/cms-service/graph/data"
	"github.com/HuanLiu-hotstar/search-federation-graphql/cms-service/graph/generated"
	"github.com/HuanLiu-hotstar/search-federation-graphql/cms-service/graph/model"
)

func (r *contentResolver) Result(ctx context.Context, obj *model.Content) ([]model.SearchResultBase, error) {
	id := obj.ID
	if strings.HasPrefix(id, "match") {
		return data.FetchScoreData(id)
	}
	return data.FetchCmsData(id)
}

// Content returns generated.ContentResolver implementation.
func (r *Resolver) Content() generated.ContentResolver { return &contentResolver{r} }

type contentResolver struct{ *Resolver }
