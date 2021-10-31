package domain

import (
	journeyDomain "citywalker/pkg/journey/domain"
	scheduleDomain "citywalker/pkg/schedule/domain"
)

type Itinerary struct {
	City string `validate:"required,lowercase" json:"city" bson:"city"`
	scheduleDomain.Schedule
	CustomEntryLocations  map[string]bool `validate:"dive,unique" json:"customEntryLocations" bson:"customEntryLocations"`
	CustomVisitLocations  []string        `validate:"dive,unique" json:"customVisitLocations" bson:"customVisitLocations"`
	Preferences           []string        `validate:"dive,unique" json:"preferences" bson:"preferences"`
	journeyDomain.Journey `validate:"required" json:"journey" bson:"journey"`
}

func (i *Itinerary) Initialize() {
	// Initialize Schedule
	i.Schedule.Initialize()

	// Initialize Journey
	i.Journey.Initialize(&i.TravelTime)
}
