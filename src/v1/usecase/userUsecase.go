package usecase

import (
	"errors"

	h "user-service/src/helper"
	"user-service/src/v1/model"
	"user-service/src/v1/repository"
)

func Registration(user *model.User) error {
	user.Password = h.HashPassword([]byte(user.Password))

	if err := repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func Login(userLogin model.UserLogin) error {
	var (
		user model.User
		err  error
	)

	user, err = repository.FetchUserByUserName(userLogin.UserName)

	if err != nil {
		return err
	}

	if !h.ComparePasswords(user.Password, []byte(userLogin.Password)) {
		return errors.New("invalid password")
	}

	return nil
}
