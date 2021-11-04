package domain

import (
	expensesDomain "citywalker/pkg/expenses/domain"
)

type Repository interface {
	Get(travelID *string) (*Travel, error)
	Gets(userID *string) (*[]Travel, error)
	Post(travel *Travel) (*Travel, error)
	Put(travel *Travel) (*Travel, error)
	PutExpenses(user *TravelUser, travelID *string, expense *expensesDomain.Expense, exists *bool) error
	PutTravelUsers(user *TravelUser, travelID *string, exists *bool) error
	Delete(travelID *string) error
}
