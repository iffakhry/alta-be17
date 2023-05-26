package user

import "time"

type Core struct {
	Id        uint
	Name      string `validate:"required"`
	Phone     string
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
}

type UserServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
}
