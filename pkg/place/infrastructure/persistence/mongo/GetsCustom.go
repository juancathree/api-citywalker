package persistence

import (
	"citywalker/pkg/place/domain"
	"context"
	"fmt"
)

func (mo *MongoRepository) GetsCustom(city *string, pref, cuV *[]string, lng *string) (*[]domain.Place, error) {
	// Create the filter
	filter := mo.CreateFilter(city, pref, cuV, lng)

	// Get from database and decode
	places := make([]domain.Place, 0, 20)
	cursor, err := mo.Collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
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
