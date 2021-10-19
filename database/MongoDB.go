package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connect DB to MongoDB database
func MongoDB() {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Defining URI to connect DB, getting environment values
	user, pass, database := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE")
	uri := fmt.Sprintf("mongodb://%v:%v@mongo.citywalker.app/%v?retryWrites=true&w=majority", user, pass, database)

	// Connecting to DB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(os.Getenv("DB_COLLECTION"))
}
