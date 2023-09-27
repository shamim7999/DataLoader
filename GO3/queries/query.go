package queries

import (
	"graphql_test/resolvers"
	"graphql_test/schema"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getBooks": &graphql.Field{
			Type:        graphql.NewList(schema.BookType),
			Description: "Returns  all books",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolvers.GetBooks,
		},

		"getAuthors": &graphql.Field{
			Type:        graphql.NewList(schema.AuthorType),
			Description: "Returns  all Authors",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolvers.GetAuthors,
		},
	},
})
