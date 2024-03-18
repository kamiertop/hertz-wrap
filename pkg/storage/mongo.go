package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitMongo init mongo driver.
func InitMongo() error {
	client, err := mongo.Connect(context.Background(), &options.ClientOptions{})
	if err != nil {
		return fmt.Errorf("connect mongo error: %w", err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		return fmt.Errorf("ping mongo error: %w", err)
	}

	return nil
}
