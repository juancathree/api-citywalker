package application

import (
	"citywalker/database"
	"citywalker/pkg/matrixCost/domain"
	persistence "citywalker/pkg/matrixCost/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
	}
}
