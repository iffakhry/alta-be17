package repositories

import (
	"be17/mvc/config"
	"be17/mvc/entities"
	"be17/mvc/middlewares"
	"be17/mvc/models"
	"errors"
)

func Login(email, password string) (entities.UserCore, string, error) {
	var user models.User
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if tx.Error != nil {
		return entities.UserCore{}, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return entities.UserCore{}, "", errors.New("login failed, email dan password salah")
	}

	token, errToken := middlewares.CreateToken(int(user.ID))
	if errToken != nil {
		return entities.UserCore{}, "", errToken
	}

	dataCore := entities.UserCore{
		Id:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return dataCore, token, nil
}
