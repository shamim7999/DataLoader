package queries

import (
	"graphql_test/resolvers"
	"graphql_test/schema"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

		"createAuthor": &graphql.Field{
			Type:        schema.AuthorType,
			Description: "Create New Author With Given Parameter",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.CreateNewAuthor,
		},
		"createBook": &graphql.Field{
			Type:        schema.BookType,
			Description: "Create New Book by Parameter",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"author_ids": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: resolvers.CreateNewBook,
		},
	},
})
