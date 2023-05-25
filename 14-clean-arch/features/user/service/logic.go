package service

import "be17/cleanarch/features/user"

type userService struct {
	userData user.UserDataInterface
}

// Create implements user.UserServiceInterface
func (*userService) Create(input user.Core) error {
	panic("unimplemented")
}

// GetAll implements user.UserServiceInterface
func (service *userService) GetAll() ([]user.Core, error) {
	data, err := service.userData.SelectAll()
	return data, err
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}
