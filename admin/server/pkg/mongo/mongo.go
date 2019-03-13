package mongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func Run() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}