package application

import (
	itineraryApplication "citywalker/pkg/itinerary/application"
	"citywalker/pkg/travel/domain"
)

func Put(travel *domain.Travel, lng *string) (*domain.Travel, error) {

	// initialize itinerary
	travel.Itinerary.Initialize()

	// Call use case create routes
	err := itineraryApplication.CreateItinerary(&travel.Itinerary, lng)
	if err != nil {
		return nil, err
	}

	modifiedTravel, err := repo.Put(travel)
	if err != nil {
		return nil, err
	}

	return modifiedTravel, nil
}
