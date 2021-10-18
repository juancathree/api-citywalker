package application

import "citywalker/pkg/user/domain"

func Get(userID *string) (*domain.User, error) {
	return repo.Get(userID)
}
