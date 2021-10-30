package application

import (
	"citywalker/database"
	"citywalker/pkg/place/domain"
	persistence "citywalker/pkg/place/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
	}
}
