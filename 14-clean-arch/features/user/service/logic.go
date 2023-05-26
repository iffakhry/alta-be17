package service

import (
	"be17/cleanarch/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Login implements user.UserServiceInterface
func (service *userService) Login(email string, password string) (user.Core, string, error) {
	if email == "" || password == "" {
		return user.Core{}, "", errors.New("error validation: nama, email, password harus diisi")
	}
	dataLogin, token, err := service.userData.Login(email, password)
	return dataLogin, token, err
}

// Create implements user.UserServiceInterface
func (service *userService) Create(input user.Core) error {
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("error validation: nama, email, password harus diisi")
	// }
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.userData.Insert(input)
	return errInsert
}

// GetAll implements user.UserServiceInterface
func (service *userService) GetAll() ([]user.Core, error) {
	data, err := service.userData.SelectAll()
	return data, err
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
