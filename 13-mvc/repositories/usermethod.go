package repositories

import (
	"be17/mvc/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func (ur UserRepo) GetAllUser() ([]models.User, error) {
	var usersData []models.User
	tx := ur.Db.Find(&usersData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}
	return usersData, nil
}

func (ur UserRepo) CreateUser(userInput models.User) error {

	tx := ur.Db.Create(&userInput) // insert into users set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
