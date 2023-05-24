package controllers

import (
	"be17/mvc/config"
	"be17/mvc/models"
	"be17/mvc/repositories"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {

	// memanggil func di repositories
	// results, err := repositories.GetAllUser()

	// memanggil method di repositories
	var userRepo = repositories.UserRepo{
		Db: config.DB,
	}
	results, err := userRepo.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error read data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success read data",
		"data":    results,
	})

}

func CreateUserController(c echo.Context) error {
	userInput := models.User{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	log.Println("input:", userInput)
	//validasi inputan dari user
	if userInput.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. nama harus diisi",
		})
	}
	err := repositories.CreateUser(userInput)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "error insert data",
			})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "error insert data",
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success insert data",
	})
}
