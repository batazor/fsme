// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type MongoFsm struct {
	Fsm      *Fsm
	ClientDB *mongo.Client
}

// Create a new MongoFsm
func (m *MongoFsm) Create(payload *Fsm) (error, *Fsm) {
	collection := m.ClientDB.Database("fsme").Collection("Fsm")

	insertResult, err := collection.InsertOne(context.TODO(), payload)
	if err != nil {
		return err, nil
	}

	payload.ID = insertResult.InsertedID.(string)

	return nil, payload
}

// Get by Id
func (m *MongoFsm) GetById(b []byte) (error, Fsm) {
	collection := m.ClientDB.Database("fsme").Collection("Fsm")
	filter := bson.D{}

	// create a value into which the result can be decoded
	var result Fsm

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err, result
	}

	return nil, result
}

// Get list
func (m *MongoFsm) GetList() (error, []*Fsm) {
	collection := m.ClientDB.Database("fsme").Collection("Fsm")

	// Pass these options to the Find method
	findOptions := options.FindOptions{}
	findOptions.SetLimit(2)
	filter := bson.D{}

	// Here's an array in which you can store the decoded documents
	var results []*Fsm

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), filter, &findOptions)
	if err != nil {
		return err, results
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Fsm
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return err, nil
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return nil, results
}

// Update
func (m *MongoFsm) Update(b []byte) (error, *mongo.UpdateResult) {
	collection := m.ClientDB.Database("fsme").Collection("Fsm")
	filter := bson.D{}

	update := bson.D{
		{"$inc", bson.D{}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err, updateResult
	}

	return nil, updateResult
}

// Delete
func (m *MongoFsm) Destroy(b []byte) (error, *mongo.DeleteResult) {
	collection := m.ClientDB.Database("fsme").Collection("Fsm")

	deleteResult, err := collection.DeleteMany(context.TODO(), nil)
	if err != nil {
		return err, nil
	}

	return nil, deleteResult
}
