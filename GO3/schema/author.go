package schema

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"graphql_test/db"
	"graphql_test/db/models"
)

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AuthorLoader = dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var authorIDs []string
	for _, key := range keys {
		authorID := key.String()
		authorIDs = append(authorIDs, authorID)
	}

	var authors []*models.Author
	filter := bson.M{"_id": bson.M{"$in": authorIDs}}
	cursor, err := db.CollectionAuthor.Find(ctx, filter)
	if err != nil {
		results := make([]*dataloader.Result, len(keys))
		for i := range results {
			results[i] = &dataloader.Result{Error: err}
		}
		return results
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var author models.Author
		if err := cursor.Decode(&author); err != nil {
			results := make([]*dataloader.Result, len(keys))
			for i := range results {
				results[i] = &dataloader.Result{Error: err}
			}
			return results
		}
		authors = append(authors, &author)
	}

	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		authorID := key.String()
		var matchingAuthor *models.Author
		for _, author := range authors {
			if author.ID.Hex() == authorID {
				matchingAuthor = author
				break
			}
		}
		if matchingAuthor != nil {
			results[i] = &dataloader.Result{Data: matchingAuthor}
		} else {
			results[i] = &dataloader.Result{Error: fmt.Errorf("Author not found: %s", authorID)}
		}
	}

	return results
})
