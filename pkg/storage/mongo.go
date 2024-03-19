package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// InitMongo init mongo driver.
func InitMongo() (err error) {
	client, err = mongo.Connect(context.Background(), &options.ClientOptions{})
	if err != nil {
		return fmt.Errorf("connect mongo error: %w", err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		return fmt.Errorf("ping mongo error: %w", err)
	}

	return nil
}

// Client return mongo client.
func Client() *mongo.Client {
	return client
}

// Database return database by name.
func Database(name string) *mongo.Database {
	return client.Database(name)
}

// CloseMongo disconnect mongo.
func CloseMongo(ctx context.Context) error {
	return client.Disconnect(ctx)
}
