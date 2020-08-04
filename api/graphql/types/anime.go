package types

import "github.com/graphql-go/graphql"

var AnimeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Anime",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Float,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"episodeNo": &graphql.Field{
			Type: graphql.Int,
		},
	},
})