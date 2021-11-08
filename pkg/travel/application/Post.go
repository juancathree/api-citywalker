package application

import (
	itineraryApplication "citywalker/pkg/itinerary/application"
	"citywalker/pkg/travel/domain"
	userApplication "citywalker/pkg/user/application"
)

func Post(travel *domain.Travel, userID string, lng *string) (*domain.Travel, error) {
	// get user
	user, err := userRepo.Get(&userID)
	if err != nil {
		return nil, err
	}

	// Initialize travel
	travel.Initialize(user)

	// Call use case create routes
	err = itineraryApplication.CreateItinerary(&travel.Itinerary, lng)
	if err != nil {
		return nil, err
	}

	newTravel, err := repo.Post(travel)
	if err != nil {
		return nil, err
	}

	err = userApplication.PostTravel(&userID, &newTravel.ID)
	if err != nil {
		return nil, err
	}

	return newTravel, nil
}
