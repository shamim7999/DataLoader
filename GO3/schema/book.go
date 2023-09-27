package schema

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	db "graphql_test/db"
	"graphql_test/db/models"
	queries "graphql_test/db/queries"
	"graphql_test/domain"
)

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author_ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"authors": &graphql.Field{
			Type: graphql.NewList(AuthorType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source := p.Source.(*domain.Book)
				var oIds []interface{}
				for _, item := range source.AuthorIds {
					oid, err := primitive.ObjectIDFromHex(item)
					if err == nil {
						oIds = append(oIds, oid)
					}
				}
				x, err := queries.GetDataFromAuthorCollection(bson.M{"_id": bson.M{"$in": oIds}})
				return x, err
			},
		},
	},
})

var BookLoader = dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var bookIDs []string
	for _, key := range keys {
		bookID := key.String()
		bookIDs = append(bookIDs, bookID)
	}

	var books []*models.Book
	filter := bson.M{"_id": bson.M{"$in": bookIDs}}
	cursor, err := db.CollectionBook.Find(ctx, filter)
	if err != nil {
		results := make([]*dataloader.Result, len(keys))
		for i := range results {
			results[i] = &dataloader.Result{Error: err}
		}
		return results
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			results := make([]*dataloader.Result, len(keys))
			for i := range results {
				results[i] = &dataloader.Result{Error: err}
			}
			return results
		}
		books = append(books, &book)
	}

	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		bookID := key.String()
		var matchingBook *models.Book
		for _, book := range books {
			if book.ID.Hex() == bookID {
				matchingBook = book
				break
			}
		}
		if matchingBook != nil {
			results[i] = &dataloader.Result{Data: matchingBook}
		} else {
			results[i] = &dataloader.Result{Error: fmt.Errorf("Book not found: %s", bookID)}
		}
	}

	return results
})
