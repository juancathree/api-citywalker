package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) UpdateDocument(travelID *string, update *primitive.M) error {
	// Create the filter
	filter, err := mo.CreateFilterID(travelID)
	if err != nil {
		return err
	}

	// Update the document in MongoDB
	_, err = mo.Collection.UpdateOne(context.Background(), filter, *update)
	if err != nil {
		return err
	}

	return nil
}
