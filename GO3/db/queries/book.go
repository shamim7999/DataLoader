package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"graphql_test/db"
	"graphql_test/db/models"
	"graphql_test/domain"
	"log"
)

func GetDataFromBookCollection(filter bson.M) ([]*domain.Book, error) {
	cursor, err := db.CollectionBook.Find(db.Ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cursor, db.Ctx)

	var books []*domain.Book

	for cursor.Next(db.Ctx) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			log.Println(err)
		}
		books = append(books, &domain.Book{
			ID:        book.ID.Hex(),
			Title:     book.Title,
			AuthorIds: book.AuthorIds,
		})
	}
	return books, nil
}

func InsertBook(book *models.Book) (*domain.Book, error) {
	if _, err := db.CollectionBook.InsertOne(db.Ctx, book); err != nil {
		return nil, err
	}
	return &domain.Book{
		ID:        book.ID.Hex(),
		Title:     book.Title,
		AuthorIds: book.AuthorIds,
	}, nil
}
