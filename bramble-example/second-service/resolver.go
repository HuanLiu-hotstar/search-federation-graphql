package main

import (
	"context"
	"strings"

	"github.com/vektah/gqlparser/v2/formatter"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Service(ctx context.Context) (*Service, error) {
	s := new(strings.Builder)
	f := formatter.NewFormatter(s)
	// parsedSchema is in the generated code
	f.FormatSchema(parsedSchema)

	service := Service{
		Name:    "gqlgen-service",
		Version: "0.1.0",
		Schema:  s.String(),
	}
	return &service, nil
}

func (r *queryResolver) Foo(ctx context.Context, id string) (*Foo, error) {
	f := fetchFooData(id)
	return &f, nil
}

func fetchFooData(id string) Foo {
	return Foo{
		ID:    id,
		Hello: "world_" + id,
		World: "hello_" + id,
	}
}
