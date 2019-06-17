package mongo

import (
	"context"
	"github.com/batazor/fsme/admin/server/pkg/model"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (db DB) List() ([]*model.FSM, error) {
	collection := db.DB.Collection(viper.GetString("MONGODB_COLLECTION"))

	// Pass these options to the Find method
	findOptions := options.FindOptions{}
	filter := bson.D{}

	// Here's an array in which you can store the decoded documents
	results := make([]*model.FSM, 0)

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), filter, &findOptions)
	if err != nil {
		return results, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem *model.FSM
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}

func (db *DB) Get(id string) (*model.FSM, error) {
	collection := db.DB.Collection(viper.GetString("MONGODB_COLLECTION"))
	idFSM, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", idFSM}}

	// Passing nil as the filter matches all documents in the collection
	response := collection.FindOne(context.TODO(), filter)

	if response.Err() != nil {
		return nil, response.Err()
	}

	var fsm model.FSM
	response.Decode(&fsm)

	return &fsm, nil
}

func (db *DB) Add(fsm model.FSM) (*primitive.ObjectID, error) {
	collection := db.DB.Collection(viper.GetString("MONGODB_COLLECTION"))

	insertResult, err := collection.InsertOne(context.TODO(), fsm)
	if err != nil {
		return nil, err
	}

	objectId := insertResult.InsertedID.(primitive.ObjectID)
	return &objectId, err
}

func (db *DB) Update(fsm model.FSM) (model.FSM, error) {
	collection := db.DB.Collection(viper.GetString("MONGODB_COLLECTION"))

	filter := bson.D{{"_id", fsm.Id}}

	// Passing nil as the filter matches all documents in the collection
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": fsm})
	return fsm, err
}

func (db *DB) Delete(idFSM string) error {
	collection := db.DB.Collection(viper.GetString("MONGODB_COLLECTION"))

	id, err := primitive.ObjectIDFromHex(idFSM)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}

	// Passing nil as the filter matches all documents in the collection
	_, err = collection.DeleteOne(context.TODO(), filter)

	return err
}
