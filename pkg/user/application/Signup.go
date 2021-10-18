package application

import (
	"citywalker/pkg/user/domain"
	"citywalker/shared"
	"fmt"
)

func Signup(user *domain.User) (*domain.User, error) {

	err := shared.Validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Set password
	user.SetPassword(&user.Password)

	return repo.Post(user)
}
