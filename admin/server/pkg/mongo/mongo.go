package mongo

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
)

func init() {
	viper.SetDefault("MONGO_URI", "mongodb://127.0.0.1:27017")
	viper.SetDefault("MONGO_DATABASE", "fsme")
	viper.SetDefault("MONGO_COLLECTION", "fsm")
	viper.SetDefault("MONGO_SINGLE", false)
}

var Cfg Config

func Run() (*mongo.Client, error) {
	var err error

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	mongoCfg := options.Client().ApplyURI(viper.GetString("MONGO_URI"))
	mongoCfg.SetDirect(viper.GetBool("MONGO_SINGLE")) // for work with single node
	Cfg.Client, err = mongo.Connect(context.TODO(), mongoCfg)

	if err != nil {
		return Cfg.Client, err
	}

	// Check the connection
	err = Cfg.Client.Ping(context.TODO(), nil)

	if err != nil {
		return Cfg.Client, err
	}

	logger.Info("Run mongodb")

	return Cfg.Client, err
}

func (c Config) List() ([]*FSM, error) {
	collection := c.Client.Database(viper.GetString("MONGO_DATABASE")).Collection(viper.GetString("MONGO_COLLECTION"))

	// Pass these options to the Find method
	findOptions := options.FindOptions{}
	filter := bson.D{}

	// Here's an array in which you can store the decoded documents
	var results []*FSM

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), filter, &findOptions)
	if err != nil {
		return results, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem *FSM
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

	//reate new FSM
	//newFSM, err := fsm.New()
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Add rule
	//newFSM.AddStateTransitionRules("a", "b", "c")
	//newFSM.AddStateTransitionRules("b", "d", "e")
	//newFSM.AddStateTransitionRules("c", "k")
	//newFSM.AddStateTransitionRules("d", "a")
	//newFSM.AddStateTransitionRules("e", "k")
	//newFSM.AddStateTransitionRules("k")
	//
	//// Add Events
	//newFSM.AddEvent("start", "a")
	//newFSM.AddEvent("to b", "b")
	//newFSM.AddEvent("to d", "d")
	//
	//// Init State
	//err = newFSM.SetStateTransition("a")
	//if err != nil {
	//	return nil, err
	//}
	//
	//state = append(state, newFSM)
	//
	//return stat
	//// Ce, nil
}

func (c Config) Add(fsm FSM) (primitive.ObjectID, error) {
	collection := c.Client.Database(viper.GetString("MONGO_DATABASE")).Collection(viper.GetString("MONGO_COLLECTION"))
	insertResult, err := collection.InsertOne(context.TODO(), fsm)
	return insertResult.InsertedID.(primitive.ObjectID), err
}

func (c Config) Update(fsm FSM) (FSM, error) {
	collection := c.Client.Database(viper.GetString("MONGO_DATABASE")).Collection(viper.GetString("MONGO_COLLECTION"))

	filter := bson.D{{"_id", fsm.Id}}

	// Passing nil as the filter matches all documents in the collection
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": fsm})
	return fsm, err
}

func (c Config) Delete(idFSM string) error {
	collection := c.Client.Database(viper.GetString("MONGO_DATABASE")).Collection(viper.GetString("MONGO_COLLECTION"))

	id, err := primitive.ObjectIDFromHex(idFSM)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}

	// Passing nil as the filter matches all documents in the collection
	_, err = collection.DeleteOne(context.TODO(), filter)

	return err
}
