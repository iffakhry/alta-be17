package data

import (
	"be17/cleanarch/features/user"

	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	// ID        string `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Name     string
	Phone    string `gorm:"unique"`
	Email    string `gorm:"unique" `
	Password string
}

// mapping dari core ke model
func CoreToModel(dataCore user.Core) User {
	return User{
		Name:     dataCore.Name,
		Phone:    dataCore.Phone,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
}
