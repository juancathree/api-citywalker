package application

import (
	itineraryDomain "citywalker/pkg/itinerary/domain"
	journeyApplication "citywalker/pkg/journey/application"
	placeDomain "citywalker/pkg/place/domain"
	"sort"
)

func CreateItinerary(itinerary *itineraryDomain.Itinerary, lng *string) error {

	// Get places
	places, err := placeRepo.GetsCustom(&itinerary.City, &itinerary.Preferences, &itinerary.CustomVisitLocations, lng)
	if err != nil {
		return err
	}

	// Get matrix costs
	matrixCost, err := matrixRepo.Get(&itinerary.City)
	if err != nil {
		return err
	}

	// Sort places by priority
	sort.Sort(placeDomain.ByPriority(*places))

	// Get places to visit
	visit := itinerary.GetVisitPlaces(places, &matrixCost.AverageCost, lng)

	// getting clusters
	var cluster Cluster
	clusters := cluster.CreateCluster(visit, matrixCost)

	// getting journeys
	journey, err := journeyApplication.CreateJourney(clusters, places, matrixCost, lng)
	if err != nil {
		return err
	}

	itinerary.Journey = *CreateItineraryJourney(itinerary, journey, places, matrixCost, lng)

	return nil
}
