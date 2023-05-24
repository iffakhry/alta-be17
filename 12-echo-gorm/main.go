package main

/*
go mod init namaproject

go get -u github.com/labstack/echo/v4
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

// struct book

// database connection
func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_raw?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	// auto migrate untuk table book
}

func init() {
	fmt.Println("ini init")
}

func main() {
	fmt.Println("ini main")
	InitDB()
	InitialMigration()
	fmt.Println("program is running")

	// init echo instance
	e := echo.New()
	// define route / endpoint
	e.POST("/users", CreateUserController)
	e.GET("/users", GetAllUserController)

	// e.GET("/users/:id", GetUserByIdController)
	// e.PUT("/users/:id", UpdateUserByIdController)
	// e.DELETE("/users/:id", DeleteUserByIdController)

	// e.POST("/books", CreateUserController)
	// e.GET("/books", GetAllUserController)

	// e.GET("/books/:id", GetUserByIdController)
	// e.PUT("/books/:id", UpdateUserByIdController)
	// e.DELETE("/books/:id", DeleteUserByIdController)

	// start server di port
	e.Logger.Fatal(e.Start(":8080"))

}

// func untuk menambah/insert data user
func CreateUserController(c echo.Context) error {
	userInput := User{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	log.Println("input:", userInput)
	// c.Request().FormFile("image")
	// proses insert data
	tx := DB.Create(&userInput) // insert into users set name = .....
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success insert data",
	})
}

// membaca seluruh data user
func GetAllUserController(c echo.Context) error {
	var usersData []User
	tx := DB.Find(&usersData) // select * from users
	if tx.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error read data",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success read data",
		"data":    usersData,
	})
}
