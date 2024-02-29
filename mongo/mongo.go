package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, clientURI string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(clientURI)
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		return nil, err
	}

	return client, nil
}
