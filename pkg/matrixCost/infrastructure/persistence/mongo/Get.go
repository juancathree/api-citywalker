package persistence

import (
	"citywalker/pkg/matrixCost/domain"
	"context"
)

func (mo *MongoRepository) Get(city *string) (*domain.MatrixCost, error) {
	// Create the filter
	filter := mo.CreateFilterID(city)

	// Get from database and decode
	var matrixCost domain.MatrixCost
	err := mo.Collection.FindOne(context.Background(), filter).Decode(&matrixCost)
	if err != nil {
		return nil, err
	}

	return &matrixCost, nil
}
