package application

import (
	"citywalker/database"
	matrixDomain "citywalker/pkg/matrixCost/domain"
	matrixPersistence "citywalker/pkg/matrixCost/infrastructure/persistence/mongo"
	placeDomain "citywalker/pkg/place/domain"
	placePersistence "citywalker/pkg/place/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var placeRepo placeDomain.Repository
var matrixRepo matrixDomain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		placeRepo = placePersistence.NewMongoRepository()
		matrixRepo = matrixPersistence.NewMongoRepository()
	}
}
