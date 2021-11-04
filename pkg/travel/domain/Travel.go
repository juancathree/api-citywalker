package domain

import (
	expensesDomain "citywalker/pkg/expenses/domain"
	itineraryDomain "citywalker/pkg/itinerary/domain"
	userDomain "citywalker/pkg/user/domain"
)

type Travel struct {
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
	itineraryDomain.Itinerary
	expensesDomain.Expenses
	TravelUsers []TravelUser `validate:"required,dive,required,unique" json:"travelUsers" bson:"travelUsers"`
}

type TravelUser struct {
	Nickname string `validate:"required" json:"nickname" bson:"nickname"`
	Email    string `validate:"required,email" json:"email" bson:"email"`
}

func (t *Travel) Initialize(user *userDomain.User) {
	// Create user struct
	u := TravelUser{
		Nickname: user.Nickname,
		Email:    user.Email,
	}

	// Initialize TravelUsers
	t.TravelUsers = append(t.TravelUsers, u)

	// Initialize Expenses
	t.Expenses.Initialize(&u.Nickname)

	// Initialize Itinerary
	t.Itinerary.Initialize()
}
