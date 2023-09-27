package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database         *mongo.Database
	CollectionBook   *mongo.Collection
	CollectionAuthor *mongo.Collection
	ClientOptions    *options.ClientOptions
	Client           *mongo.Client
	Err              error
	Ctx              context.Context
)
