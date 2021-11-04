package persistence

import (
	expensesDomain "citywalker/pkg/expenses/domain"
	"citywalker/pkg/travel/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PutTravelUsers(user *domain.TravelUser, travelID *string, exists *bool) error {
	// Create the update
	var update bson.M

	if *exists {
		update = bson.M{"$pull": bson.M{"travelUsers": *user}, "$unset": bson.M{"expenses." + *&user.Nickname: ""}}
	} else {
		expense := []expensesDomain.Expense{}
		update = bson.M{"$push": bson.M{"travelUsers": *user}, "$set": bson.M{"expenses." + *&user.Nickname: expense}}
	}

	return mo.UpdateDocument(travelID, &update)
}
