package persistence

import (
	"citywalker/pkg/place/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mo *MongoRepository) Gets(city *string, lng *string) (*[]domain.Place, error) {
	// Create the filter
	filter := mo.CreateFilterID(city, lng)

	// Create projection
	projection := mo.CreateProjection(lng)

	// Get from database and decode
	places := make([]domain.Place, 0, 30)
	cursor, err := mo.Collection.Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var place domain.Place
		err := cursor.Decode(&place)
		if err != nil {
			return nil, err
		}

		places = append(places, place)
	}

	return &places, nil
}
