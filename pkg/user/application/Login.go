package application

import (
	"citywalker/pkg/user/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Login(user *domain.User) error {
	// Get User from database
	userRepo, err := repo.Get(&user.Email)
	if err != nil {
		return err
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(userRepo.Password), []byte(user.Password))
	if err != nil {
		return fmt.Errorf("the passwords are not equals")
	}

	return nil
}
