module github.com/francesc2509/go-graphql-example/api/graphql/fields

go 1.14

require (
  github.com/graphql-go/graphql v0.7.9
  github.com/francesc2509/go-graphql-example/entities v0.0.0
  github.com/francesc2509/go-graphql-example/api/services v0.0.0
  github.com/francesc2509/go-graphql-example/api/graphql/types v0.0.0
)

replace (
  github.com/francesc2509/go-graphql-example/entities => ../../../entities
  github.com/francesc2509/go-graphql-example/api/services => ../../services
  github.com/francesc2509/go-graphql-example/api/graphql/types => ../types
)