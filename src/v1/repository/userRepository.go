package repository

import (
	"user-service/src/v1/model"
)

func CreateUser(user *model.User) error {
	db := getDB()

	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func FetchUserByUserName(email string) (model.User, error) {
	var user model.User

	db := getDB()

	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
