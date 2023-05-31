package data

import (
	"be17/cleanarch/app/middlewares"
	"be17/cleanarch/features/user"
	"be17/cleanarch/helper"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Login implements user.UserDataInterface
func (repo *userQuery) Login(email string, password string) (user.Core, string, error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm) // select * from users limit 1
	if tx.Error != nil {
		return user.Core{}, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return user.Core{}, "", errors.New("login failed, email dan password salah")
	}

	checkPassword := helper.CheckPasswordHash(password, userGorm.Password)
	if !checkPassword {
		return user.Core{}, "", errors.New("login failed, password salah")
	}

	token, errToken := middlewares.CreateToken(int(userGorm.ID))
	if errToken != nil {
		return user.Core{}, "", errToken
	}

	dataCore := user.Core{
		Id:        userGorm.ID,
		Name:      userGorm.Name,
		Phone:     userGorm.Phone,
		Email:     userGorm.Email,
		Password:  userGorm.Password,
		CreatedAt: userGorm.CreatedAt,
		UpdatedAt: userGorm.UpdatedAt,
	}
	return dataCore, token, nil
}

// Insert implements user.UserDataInterface
func (repo *userQuery) Insert(input user.Core) error {
	// mapping dari struct entities core ke gorm model
	// userInputGorm := User{
	// 	Name:     input.Name,
	// 	Phone:    input.Phone,
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// }
	hashedPassword, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	userInputGorm := CoreToModel(input)
	userInputGorm.Password = hashedPassword

	tx := repo.db.Create(&userInputGorm) // insert into users set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

// SelectAll implements user.UserDataInterface
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersData []User
	tx := repo.db.Find(&usersData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(usersData)
	// mapping dari struct gorm model ke struct entities core
	var usersCoreAll []user.Core
	for _, value := range usersData {
		var userCore = user.Core{
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
