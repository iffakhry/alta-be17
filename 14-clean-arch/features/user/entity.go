package user

import "time"

type Core struct {
	Id        uint
	Name      string
	Phone     string
	Email     string
	Password  string
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
