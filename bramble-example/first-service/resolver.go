package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
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
	foo := fetchFooData(id)
	return &foo, nil
}

func (r *queryResolver) Bar(ctx context.Context, id string) (*Bar, error) {
	bar := &Bar{
		ID:   "bar_" + id,
		Name: "name_bar_" + id,
	}
	log.Printf("call bar %+v", bar)
	return bar, nil
}
func (r *queryResolver) RandomFoo(ctx context.Context) (*Foo, error) {
	id := strconv.Itoa(rand.Intn(100))
	foo := Foo{
		ID:     id,
		Gqlgen: true,
		Bar: &Bar{
			ID:   "bar-" + id,
			Name: "bar",
		},
	}
	return &foo, nil
}

func fetchFooData(id string) Foo {
	return Foo{
		ID:     id,
		Gqlgen: true,
		Bar: &Bar{
			ID: "bar-" + id,
		},
	}
}
