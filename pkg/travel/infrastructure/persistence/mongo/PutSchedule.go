package persistence

import (
	scheduleDomain "citywalker/pkg/schedule/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PutSchedule(schedule *scheduleDomain.Schedule, travelID *string) error {
	// Create the update
	update := bson.M{
		"$set": bson.M{
			"startDay":           (*schedule).StartDay,
			"endDay":             (*schedule).EndDay,
			"itineraryStartTime": (*schedule).ItineraryStartTime,
			"itineraryEndTime":   (*schedule).ItineraryEndTime,
			"travelTime":         (*schedule).TravelTime,
		},
	}

	return mo.UpdateDocument(travelID, &update)
}
