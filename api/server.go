package api

import (
	"net/http"

	"github.com/francesc2509/go-graphql-example/api/graphql"
	servergo "github.com/francesc2509/http-wrapper"
)

func Start() {

	router := servergo.New()

	router.HandleFunc("/graphql", graphql.Init())

	http.ListenAndServe(":12345", router)
}
