package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Client, error) {

	return nil, nil
}
