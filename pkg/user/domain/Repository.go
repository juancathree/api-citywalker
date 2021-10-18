package domain

type Repository interface {
	Get(userID *string) (*User, error)
	Post(user *User) (*User, error)
	PostTravel(userID *string, travelID *string) error
}
