package persistence

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilter(city *string, pref, cuV *[]string) *primitive.M {
	// Create a filter for MongoDB query
	filter := bson.M{
		"city.en":     strings.Title(*city),
		"category.en": bson.M{"$nin": *pref},
		"name.en":     bson.M{"$nin": *cuV},
	}

	return &filter
}
