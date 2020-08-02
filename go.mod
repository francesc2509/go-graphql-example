module github.com/francesc2509/go-graphql-example

go 1.14

require (
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
  github.com/francesc2509/go-graphql-example/api/services v0.0.0
  github.com/francesc2509/go-graphql-example/api/graphql v0.0.0
  github.com/francesc2509/go-graphql-example/api/graphql/fields v0.0.0
  github.com/francesc2509/go-graphql-example/api/graphql/types v0.0.0
  github.com/francesc2509/go-graphql-example/entities v0.0.0
  github.com/francesc2509/go-graphql-example/api v0.0.0
)

replace (
  github.com/francesc2509/go-graphql-example/api/services => ./api/services
  github.com/francesc2509/go-graphql-example/api/graphql => ./api/graphql
  github.com/francesc2509/go-graphql-example/api/graphql/fields => ./api/graphql/fields
  github.com/francesc2509/go-graphql-example/api/graphql/types => ./api/graphql/types
  github.com/francesc2509/go-graphql-example/entities => ./entities
  github.com/francesc2509/go-graphql-example/api => ./api
)