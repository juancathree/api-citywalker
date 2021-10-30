package application

import "citywalker/pkg/place/domain"

func Gets(city *string, lng *string) (*[]domain.Place, error) {
	return repo.Gets(city, lng)
}
