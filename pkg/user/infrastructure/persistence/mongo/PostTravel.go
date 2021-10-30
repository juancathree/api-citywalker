package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PostTravel(userID *string, travelID *string) error {
	// Create the filter
	filter := mo.CreateFilterID(userID)

	// Create the update
	update := bson.M{"$push": bson.M{"travels": *travelID}}

	// Inserting user to database
	_, err := mo.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
