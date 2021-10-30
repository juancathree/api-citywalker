package persistence

import (
	"citywalker/pkg/user/domain"
	"context"
)

func (mo *MongoRepository) Post(user *domain.User) error {
	// Inserting user to database
	_, err := mo.Collection.InsertOne(context.Background(), *user)
	if err != nil {
		return err
	}

	return nil
}
