package persistence

import (
	expensesDomain "citywalker/pkg/expenses/domain"
	"citywalker/pkg/travel/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PutExpenses(user *domain.TravelUser, travelID *string, expense *expensesDomain.Expense, exists *bool) error {
	// Create the update
	var update bson.M

	if *exists {
		update = bson.M{"$pull": bson.M{"Expenses." + *&user.Nickname: *expense}}
	} else {
		update = bson.M{"$push": bson.M{"Expenses." + *&user.Nickname: *expense}}
	}

	return mo.UpdateDocument(travelID, &update)
}
