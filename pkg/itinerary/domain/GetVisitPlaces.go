package domain

import (
	placeDomain "citywalker/pkg/place/domain"
)

func (i *Itinerary) GetVisitPlaces(places *[]placeDomain.Place, averageCost *float64, lng *string) *[]placeDomain.Place {
	// Get total hours
	hours := i.GetTotalHours()

	visit := make([]placeDomain.Place, 0, len(*places))

	for _, place := range *places {
		// check if it is in custom entry customEntryLocations
		if i.CustomEntryLocations[place.Name[*lng]] {
			timeVisit := place.Visit.All + *averageCost
			if *hours-timeVisit > 0 {
				*hours -= timeVisit
				visit = append(visit, place)
			} else {
				break
			}
			// if isn't in customEntryLocations
		} else if place.Visit.Outside != 0 {
			timeVisit := place.Visit.Outside + *averageCost
			if *hours-timeVisit > 0 {
				*hours -= timeVisit
				visit = append(visit, place)
			} else {
				break
			}
		}
	}

	return &visit
}
