package persistence

import (
	"context"
)

func (mo *MongoRepository) Delete(travelID *string) error {
	// Create the filter
	filter, err := mo.CreateFilterID(travelID)
	if err != nil {
		return err
	}

	// Delete the document in MongoDB
	_, err = mo.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
