package persistence

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilter(city *string, pref, cuV *[]string, lng *string) *primitive.M {
	// Create a filter for MongoDB query
	filter := bson.M{
		"city." + (*lng):     strings.Title(*city),
		"category." + (*lng): bson.M{"$nin": *pref},
		"name." + (*lng):     bson.M{"$nin": *cuV},
	}

	return &filter
}
