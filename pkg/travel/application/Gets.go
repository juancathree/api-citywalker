package application

import "citywalker/pkg/travel/domain"

func Gets(userID *string) (*[]domain.Travel, error) {
	return repo.Gets(userID)
}
