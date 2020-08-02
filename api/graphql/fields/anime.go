package fields

import (
	"github.com/francesc2509/go-graphql-example/api/graphql/types"
	"github.com/francesc2509/go-graphql-example/api/services"
	"github.com/francesc2509/go-graphql-example/entities"
	"github.com/graphql-go/graphql"
)

// HandleQuery adds the graphql fields corresponding to
// artist queries
func HandleAnimeQuery(fields *graphql.Fields) {
	f := (*fields)
	f["animeGet"] = get()
}

// HandleMutation adds the graphql fields corresponding to
// artist mutations
func HandleAnimeMutation(fields *graphql.Fields) {
	f := (*fields)
	f["animeCreate"] = create()
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
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			artist := &entities.Anime{}
			artist.Title = params.Args["title"].(string)
			artist.EpisodeNo = params.Args["episodeNo"].(uint)
			return services.Anime.Create(artist)
		},
	}
}