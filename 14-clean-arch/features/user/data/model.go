package data

import (
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
