// Code generated by go-swagger; DO NOT EDIT.

package models

import "github.com/mongodb/mongo-go-driver/mongo"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Create a new Fsm
func (m *Fsm) Create(b []byte) (error, result mongo.InsertOneResult) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	insertResult, err := collection("namePackage").InsertOne(context.TODO(), b)
	if err != nil {
		return err, nil
	}

	return nil, insertResult.InsertedID
}

// Get by Id
func (m *Fsm) GetById(b []byte) (error, Fsm) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	// create a value into which the result can be decoded
	var result Fsm

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err, nil
	}

	return nil, result
}

// Get list
func (m *Fsm) GetList(b []byte) (error, Fsm) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	// Pass these options to the Find method
	options := options.Find()
	options.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*Fsm

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), nil, options)
	if err != nil {
		return err, nil
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

	return nil, result
}

// Update
func (m *Fsm) Update(b []byte) (error, mongo.UpdateResult) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")
	filter := bson.D{}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err, nil
	}

	return nil, updateResult
}

// Delete
func (m *Fsm) Destroy(b []byte) (error, result mongo.DeleteResult) {
	collection := m.ClientDB.Database("nameService").Collection("namePackage")

	deleteResult, err := collection.DeleteMany(context.TODO(), nil)
	if err != nil {
		return err, nil
	}

	return nil, deleteResult
}
