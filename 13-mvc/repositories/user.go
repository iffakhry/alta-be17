package repositories

import (
	"be17/mvc/config"
	"be17/mvc/models"
	"errors"
)

func GetAllUser() ([]models.User, error) {
	var usersData []models.User
	tx := config.DB.Find(&usersData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}
	return usersData, nil
}

func CreateUser(userInput models.User) error {

	tx := config.DB.Create(&userInput) // insert into users set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
