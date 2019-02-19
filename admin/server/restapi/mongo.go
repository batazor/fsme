package restapi

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	"context"
)

type Mongo struct {
	client *mongo.Client
}

func (m *Mongo) Connect() error {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.client, err = mongo.Connect(ctx, "mongodb://localhost:27017")

	return err
}

func (m )