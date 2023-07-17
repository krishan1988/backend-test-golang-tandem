package mongo

import (
	"context"

	"github.com/KryptoKnight/backend-test-golang/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

// Connect create a mongo database instance.
func Connect(ctx context.Context, cfg config.Database) (err error) {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	db = client.Database(cfg.Name)
	return nil
}

// GetDB return the database.
func GetDB() *mongo.Database {
	return db
}

// Close mongodb client.
func Close(ctx context.Context) {
	client.Disconnect(ctx)
}
