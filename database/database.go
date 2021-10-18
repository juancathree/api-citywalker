package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var DB interface{}

// InitDatabase
func DefineDB() {
	switch os.Getenv("DATABASE") {
	case "MONGODB":
		var db *mongo.Database
		DB = db
	}
}
