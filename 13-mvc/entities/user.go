package entities

import "time"

type UserCore struct {
	Id        uint      `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Phone     string    `json:"phone" form:"phone"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type UserResponse struct {
	Id        uint      `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
