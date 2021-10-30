package domain

import (
	"time"

	"github.com/go-playground/validator"
)

type Schedule struct {
	StartDay           time.Time `validate:"required,gt" json:"startDay" bson:"startDay"`
	EndDay             time.Time `validate:"required,gtcsfield=StartDay" json:"endDay" bson:"endDay"`
	ItineraryStartTime time.Time `validate:"required" json:"itineraryStartTime" bson:"itineraryStartTime"`
	ItineraryEndTime   time.Time `validate:"required,gtcsfield=ItineraryStartTime" json:"itineraryEndTime" bson:"itineraryEndTime"`
	TravelTime         int       `validate:"required" json:"travelTime" bson:"travelTime"`
}

// NewSchedule calculate travelTime and validate the schedule
func (s *Schedule) Initialize() error {
	// Calculate TravelTime
	s.TravelTime = int(s.EndDay.Sub(s.StartDay).Hours())/24 + 1

	// Validate schedule
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
