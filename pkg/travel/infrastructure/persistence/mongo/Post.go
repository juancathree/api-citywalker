package persistence

import (
	"citywalker/pkg/travel/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) Post(travel *domain.Travel) (*domain.Travel, error) {
	// Inserting travel to database
	result, err := mo.Collection.InsertOne(context.Background(), *travel)
	if err != nil {
		return nil, err
	}

	// Convert result to ObjectID
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}

	// Convert ObjectID to string
	travel.ID = oid.Hex()

	return travel, nil
}
