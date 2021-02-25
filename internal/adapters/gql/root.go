package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// New Graphql handler
func New() *handler.Handler {
	// Graphql
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Geotool",
		Fields: graphql.Fields{
			"matrix": Matrix(),
		},
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
