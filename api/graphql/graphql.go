package graphql

import (
	"github.com/francesc2509/go-graphql-example/api/graphql/fields"

	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type key int

// Init initializes GraphQL's root schema
func Init() http.HandlerFunc {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: fields.HandleQuery(),
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: fields.HandleMutation(),
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		h.ContextHandler(r.Context(), w, r)
	}
}
