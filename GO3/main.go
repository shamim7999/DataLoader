package main

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"graphql_test/db"
	"graphql_test/queries"
	"graphql_test/schema"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

type ReqBody struct {
	Query string `json:"query"`
}

func main() {

	db.ClientOptions = options.Client().ApplyURI(
		"mongodb://admin:secret@localhost:27017",
	).SetDirect(true)

	db.Ctx = context.Background()
	db.Client, db.Err = mongo.Connect(db.Ctx, db.ClientOptions)
	if db.Err != nil {
		log.Fatal(db.Err)
	}
	defer db.Client.Disconnect(db.Ctx)

	fmt.Println("Connected to MongoDB")

	db.Database = db.Client.Database("shamim")
	db.CollectionBook = db.Database.Collection("Book")
	db.CollectionAuthor = db.Database.Collection("Author")
	rand.Seed(time.Now().UnixNano())

	customMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "authorLoader", schema.AuthorLoader)
			ctx = context.WithValue(ctx, "bookLoader", schema.BookLoader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
	graphqlHandler := handler.New(&handler.Config{
		Schema: queries.GetRootSchema(),
		Pretty: true,
	})
	http.Handle("/graphql", customMiddleware(graphqlHandler))

	fmt.Println("Now server is running on port 8080")

	http.ListenAndServe(":8080", nil)
}
