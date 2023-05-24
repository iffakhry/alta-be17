package controllers

import (
	"be17/mvc/entities"
	"be17/mvc/helper"
	"be17/mvc/repositories"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {

	// memanggil func di repositories
	results, err := repositories.GetAllUser()

	// memanggil method di repositories
	// var userRepo = repositories.UserRepo{
	// 	Db: config.DB,
	// }
	// results, err := userRepo.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var usersResponse []entities.UserResponse
	for _, value := range results {
		usersResponse = append(usersResponse, entities.UserResponse{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", usersResponse))

}

func CreateUserController(c echo.Context) error {
	userInput := entities.UserCore{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	log.Println("input:", userInput)
	//validasi inputan dari user
	if userInput.Name == "" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error. nama harus diisi"))
	}
	err := repositories.CreateUser(userInput)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert data, row affected = 0"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}
