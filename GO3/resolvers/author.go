package resolvers

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"graphql_test/db/models"
	db "graphql_test/db/queries"
)

func GetAuthors(p graphql.ResolveParams) (interface{}, error) {
	return db.GetDataFromAuthorCollection(bson.M{})
}

func CreateNewAuthor(p graphql.ResolveParams) (interface{}, error) {
	var name string

	if val, ok := p.Args["name"].(string); ok {
		name = val
	}

	return db.InsertAuthor(&models.Author{
		ID:   primitive.NewObjectID(),
		Name: name,
	})
}
