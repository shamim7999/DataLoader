package queries

import (
	"github.com/graphql-go/graphql"
)

func GetRootSchema() *graphql.Schema {
	var rootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	})
	return &rootSchema
}
