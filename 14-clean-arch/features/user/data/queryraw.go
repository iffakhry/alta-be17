package data

import (
	"be17/cleanarch/features/user"

	"gorm.io/gorm"
)

type userQueryRaw struct {
	db *gorm.DB
}

// Login implements user.UserDataInterface
func (*userQueryRaw) Login(email string, password string) (user.Core, string, error) {
	panic("unimplemented")
}

// Insert implements user.UserDataInterface
func (*userQueryRaw) Insert(input user.Core) error {
	panic("unimplemented")
}

// SelectAll implements user.UserDataInterface
func (*userQueryRaw) SelectAll() ([]user.Core, error) {
	panic("unimplemented")
}

func NewRaw(db *gorm.DB) user.UserDataInterface {
	return &userQueryRaw{
		db: db,
	}
}
