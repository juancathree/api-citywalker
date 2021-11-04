package application

import (
	"citywalker/database"
	"citywalker/pkg/travel/domain"
	persistence "citywalker/pkg/travel/infrastructure/persistence/mongo"
	userDomain "citywalker/pkg/user/domain"
	userPersistence "citywalker/pkg/user/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository
var userRepo userDomain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
		userRepo = userPersistence.NewMongoRepository()
	}
}
