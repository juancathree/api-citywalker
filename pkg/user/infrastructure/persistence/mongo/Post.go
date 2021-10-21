package persistence

import (
	"citywalker/pkg/user/domain"
	"context"
	"fmt"
)

func (mo *MongoRepository) Post(user *domain.User) error {
	// Inserting user to database
	_, err := mo.Collection.InsertOne(context.Background(), *user)
	if err != nil {
		return fmt.Errorf("database error: couldn't be posted user to database")
	}

	return nil
}
