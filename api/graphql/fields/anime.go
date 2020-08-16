package fields

import (
	"fmt"

	"github.com/francesc2509/go-graphql-example/api/graphql/types"
	"github.com/francesc2509/go-graphql-example/api/services"
	"github.com/francesc2509/go-graphql-example/entities"
	"github.com/graphql-go/graphql"
)

// HandleQuery adds the graphql fields corresponding to
// anime queries
func HandleAnimeQuery(fields *graphql.Fields) {
	f := (*fields)
	f["animeGet"] = get()
}

// HandleMutation adds the graphql fields corresponding to
// anime mutations
func HandleAnimeMutation(fields *graphql.Fields) {
	f := (*fields)
	f["animeCreate"] = create()
	f["animeModify"] = modify()
}

func get() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.AnimeType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return services.Anime.Get()
		},
	}
}

func create() *graphql.Field {
	return &graphql.Field{
		Type: types.AnimeType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"episodeNo": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			anime := &entities.Anime{}
			anime.Title = params.Args["title"].(string)
			anime.EpisodeNo = uint(params.Args["episodeNo"].(int))
			err := services.Anime.Create(anime)

			fmt.Println(anime)
			return anime, err
		},
	}
}

func modify() *graphql.Field {
	return &graphql.Field{
		Type: types.AnimeType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"episodeNo": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			anime := &entities.Anime{}
			anime.Id = uint64(params.Args["id"].(int))
			anime.Title = params.Args["title"].(string)
			anime.EpisodeNo = uint(params.Args["episodeNo"].(int))
			anime, err := services.Anime.Modify(anime)

			return anime, err
		},
	}
}
