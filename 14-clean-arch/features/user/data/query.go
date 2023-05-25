package data

import (
	"be17/cleanarch/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Insert implements user.UserDataInterface
func (*userQuery) Insert(input user.Core) error {
	panic("unimplemented")
}

// SelectAll implements user.UserDataInterface
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersData []User
	tx := repo.db.Find(&usersData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}

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

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
