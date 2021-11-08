package application

import (
	"citywalker/pkg/travel/domain"
	userApplication "citywalker/pkg/user/application"
)

func PutUsers(userID, travelID string, exist bool) error {

	// get user
	user, err := userRepo.Get(&userID)
	if err != nil {
		return err
	}

	u := domain.TravelUser{
		Nickname: user.Nickname,
		Email:    user.Email,
	}

	err = repo.PutTravelUsers(&u, &travelID, &exist)
	if err != nil {
		return err
	}

	err = userApplication.PostTravel(&user.ID, &travelID)
	if err != nil {
		return err
	}

	return nil
}
