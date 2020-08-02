module github.com/francesc2509/go-graphql-example/api/graphql

go 1.14

require (
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
  github.com/francesc2509/go-graphql-example/api/graphql/fields v0.0.0
  github.com/francesc2509/go-graphql-example/api/graphql/types v0.0.0
  github.com/francesc2509/go-graphql-example/entities v0.0.0
)

replace (
  github.com/francesc2509/go-graphql-example/api/services => ../services
  github.com/francesc2509/go-graphql-example/api/graphql/fields => ./fields
  github.com/francesc2509/go-graphql-example/api/graphql/types => ./types
  github.com/francesc2509/go-graphql-example/entities => ../../entities
)