package repositories

import (
	"be17/mvc/entities"
	"be17/mvc/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func (ur UserRepo) GetAllUser() ([]entities.UserCore, error) {
	var usersData []models.User
	tx := ur.Db.Find(&usersData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping dari struct gorm model ke struct entities core
	var usersCoreAll []entities.UserCore
	for _, value := range usersData {
		var userCore = entities.UserCore{
			Id:        value.ID,
			Name:      value.Name,
			Phone:     value.Phone,
			Email:     value.Email,
			Password:  value.Password,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		usersCoreAll = append(usersCoreAll, userCore)
	}
	return usersCoreAll, nil
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
