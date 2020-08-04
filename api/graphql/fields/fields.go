package fields

import (
	"github.com/graphql-go/graphql"
)

// HandleQuery returns the queries graphql fields added
func HandleQuery() graphql.Fields {
	queryFields := make(graphql.Fields)

	HandleAnimeQuery(&queryFields)
	// anime.HandleQuery(&queryFields)
	// question.HandleQuery(&queryFields)

	return queryFields
}

// HandleMutation returns the mutations graphql fields added
func HandleMutation() graphql.Fields {
	mutationFields := make(graphql.Fields)

	HandleAnimeMutation(&mutationFields)
	// anime.HandleMutation(&mutationFields)
	// user.HandleMutation(&mutationFields)

	return mutationFields
}
