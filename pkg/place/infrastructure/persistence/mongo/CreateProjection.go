package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateProjection(lng *string) *primitive.D {
	// Create a filter for MongoDB query
	projection := bson.D{
		primitive.E{Key: "_id", Value: 1},
		primitive.E{Key: "city." + (*lng), Value: 1},
		primitive.E{Key: "name." + (*lng), Value: 1},
		primitive.E{Key: "location", Value: 1},
		primitive.E{Key: "information." + (*lng), Value: 1},
		primitive.E{Key: "visit", Value: 1},
		primitive.E{Key: "category." + (*lng), Value: 1},
		primitive.E{Key: "priority", Value: 1},
	}

	return &projection
}
