package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) GetCities(lng *string) (*[]string, error) {

	// Get from database and decode
	result, err := mo.Collection.Distinct(context.Background(), "city."+(*lng), bson.M{})
	if err != nil {
		return nil, err
	}

	resultString := make([]string, 0, len(result))

	for _, r := range result {
		resultString = append(resultString, r.(string))
	}

	return &resultString, nil
}
