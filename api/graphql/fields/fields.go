package fields

import (
	"github.com/graphql-go/graphql"
)

// HandleQuery returns the queries graphql fields added
func HandleQuery() graphql.Fields {
	queryFields := make(graphql.Fields)

	HandleAnimeQuery(&queryFields)
	// artist.HandleQuery(&queryFields)
	// question.HandleQuery(&queryFields)

	return queryFields
}

// HandleMutation returns the mutations graphql fields added
func HandleMutation() graphql.Fields {
	mutationFields := make(graphql.Fields)

	HandleAnimeMutation(&mutationFields)
	// artist.HandleMutation(&mutationFields)
	// user.HandleMutation(&mutationFields)

	return mutationFields
}
