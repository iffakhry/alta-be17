package models

import "gorm.io/gorm"

// struct gorm model
type User struct {
	gorm.Model
	// ID        string `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Name     string `json:"name" form:"name"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
