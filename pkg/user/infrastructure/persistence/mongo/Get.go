package persistence

import (
	"citywalker/pkg/user/domain"
	"context"
)

func (mo *MongoRepository) Get(userID *string) (*domain.User, error) {
	// Create the filter
	filter := mo.CreateFilterID(userID)

	// Get from database and decode
	var user domain.User
	err := mo.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
