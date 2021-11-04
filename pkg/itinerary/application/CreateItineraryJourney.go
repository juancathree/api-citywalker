package application

import (
	"citywalker/pkg/itinerary/domain"
	journeyDomain "citywalker/pkg/journey/domain"
	matrixDomain "citywalker/pkg/matrixCost/domain"
	placeDomain "citywalker/pkg/place/domain"
)

func CreateItineraryJourney(itinerary *domain.Itinerary, journey *[]placeDomain.Place, places *[]placeDomain.Place, matrixCost *matrixDomain.MatrixCost, lng *string) *journeyDomain.Journey {

	var j journeyDomain.Journey
	j.Initialize(&itinerary.TravelTime)

	k := 0

	for i := 0; i < itinerary.TravelTime; i++ {
		time := itinerary.Schedule.HoursOf(i)

		first := true

		for time > 0 && k < len(*journey) {

			currentPlace := (*journey)[k]

			if first {
				j[i] = append(j[i], currentPlace.Name[*lng])

				if itinerary.CustomEntryLocations[currentPlace.Name[*lng]] {
					time -= currentPlace.Visit.All
				} else {
					time -= currentPlace.Visit.Outside
				}

				first = false
				k++

			} else {

				previousPlace := (*journey)[k-1]

				if itinerary.CustomEntryLocations[currentPlace.Name[*lng]] {

					timeVisit := currentPlace.Visit.All + (*matrixCost).Costs[currentPlace.Name[*lng]][previousPlace.Name[*lng]]

					if time > timeVisit {
						time -= timeVisit
						j[i] = append(j[i], currentPlace.Name[*lng])
						k++
					} else {
						break
					}

				} else {

					timeVisit := currentPlace.Visit.Outside + (*matrixCost).Costs[currentPlace.Name[*lng]][previousPlace.Name[*lng]]

					if time > timeVisit {
						time -= timeVisit
						j[i] = append(j[i], currentPlace.Name[*lng])
						k++
					} else {
						break
					}
				}
			}
		}
	}

	return &j
}
