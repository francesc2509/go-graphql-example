package api

import (
	"net/http"
	"github.com/francesc2509/go-graphql-example/api/graphql"
)

func Start() {
	http.HandleFunc("/graphql", graphql.Init())

	http.ListenAndServe(":12345", nil)
}