package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}
