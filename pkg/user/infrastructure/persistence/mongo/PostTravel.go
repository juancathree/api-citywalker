package persistence

import (
	"context"
	"fmt"

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
		return fmt.Errorf("database error: couldn't be added travel to user")
	}

	return nil
}
