package application

import (
	"citywalker/pkg/user/domain"
	"citywalker/shared"
)

func Signup(user *domain.User) error {

	err := shared.Validate.Struct(user)
	if err != nil {
		return err
	}
	// Set password
	user.SetPassword(&user.Password)

	return repo.Post(user)
}
