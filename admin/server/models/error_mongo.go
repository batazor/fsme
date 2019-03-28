// Code generated by go-swagger; DO NOT EDIT.

package models

import (
	"context"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

type MongoError struct {
	Error    *Error
	ClientDB *mongo.Client
}

// Create a new MongoError
func (m *MongoError) Create(b []byte) (error, interface{}) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	insertResult, err := collection.InsertOne(context.TODO(), b)
	if err != nil {
		return err, nil
	}

	return nil, insertResult.InsertedID
}

// Get by Id
func (m *MongoError) GetById(b []byte) (error, Error) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")
	filter := bson.D{}

	// create a value into which the result can be decoded
	var result Error

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err, result
	}

	return nil, result
}

// Get list
func (m *MongoError) GetList(b []byte) (error, []*Error) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	// Pass these options to the Find method
	options := options.Find()
	options.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*Error

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), nil, options)
	if err != nil {
		return err, results
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Error
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
func (m *MongoError) Update(b []byte) (error, *mongo.UpdateResult) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")
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
func (m *MongoError) Destroy(b []byte) (error, *mongo.DeleteResult) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	deleteResult, err := collection.DeleteMany(context.TODO(), nil)
	if err != nil {
		return err, nil
	}

	return nil, deleteResult
}
