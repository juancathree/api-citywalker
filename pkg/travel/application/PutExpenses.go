package application

import (
	expenseDomain "citywalker/pkg/expenses/domain"
	"citywalker/pkg/travel/domain"
)

func PutExpenses(userID, travelID string, spend *expenseDomain.Expense, exist bool) error {

	// get user
	user, err := userRepo.Get(&userID)
	if err != nil {
		return err
	}

	u := domain.TravelUser{
		Nickname: user.Nickname,
		Email:    user.Email,
	}

	err = repo.PutExpenses(&u, &travelID, spend, &exist)
	if err != nil {
		return err
	}

	return nil
}
