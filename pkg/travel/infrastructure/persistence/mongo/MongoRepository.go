package persistence

import (
	"citywalker/database"
	"citywalker/pkg/travel/domain"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository() domain.Repository {
	return &MongoRepository{
		Collection: database.DB.(*mongo.Database).Collection(os.Getenv("DB_COLLECTION_TRAVEL")),
	}
}
