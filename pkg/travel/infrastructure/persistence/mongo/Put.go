package persistence

import (
	"citywalker/pkg/travel/domain"
	"context"
)

func (mo *MongoRepository) Put(travel *domain.Travel) (*domain.Travel, error) {

	// Create the filter
	filter, err := mo.CreateFilterID(&travel.ID)
	if err != nil {
		return nil, err
	}

	travel.ID = ""

	// Inserting travel to database
	_, err = mo.Collection.ReplaceOne(context.Background(), filter, *travel)
	if err != nil {
		return nil, err
	}

	return travel, nil
}
