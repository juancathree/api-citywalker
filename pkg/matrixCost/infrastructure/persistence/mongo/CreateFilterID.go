package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilterID(city *string) *primitive.M {
	// Create a filter for MongoDB query
	filter := bson.M{"city": *city}

	return &filter
}
