package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	// Set type of DB from .env
	DefineDB()

	// Connecting respective DB type to its database
	switch DB.(type) {
	case *mongo.Database:
		MongoDB()
	}
}
