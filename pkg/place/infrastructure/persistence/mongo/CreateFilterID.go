package persistence

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilterID(city *string, lng *string) *primitive.M {
	// Create a filter for MongoDB query
	filter := bson.M{"city." + (*lng): strings.Title(*city)}

	return &filter
}
