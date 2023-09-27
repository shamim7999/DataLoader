package resolvers

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"graphql_test/db/models"
	db "graphql_test/db/queries"
)

func GetBooks(p graphql.ResolveParams) (interface{}, error) {

	return db.GetDataFromBookCollection(bson.M{})
}

func CreateNewBook(p graphql.ResolveParams) (interface{}, error) {
	var (
		authorIds []string
		title     string
	)

	if val, ok := p.Args["title"].(string); ok {
		title = val
	}
	if ids, ok := p.Args["author_ids"].([]interface{}); ok {
		for _, item := range ids {
			authorIds = append(authorIds, item.(string))
		}
	}
	return db.InsertBook(&models.Book{
		ID:        primitive.NewObjectID(),
		Title:     title,
		AuthorIds: authorIds,
	})
}
