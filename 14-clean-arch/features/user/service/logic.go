package service

import (
	"be17/cleanarch/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
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
