package controllers

import (
	"be17/mvc/entities"
	"be17/mvc/helper"
	"be17/mvc/repositories"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	loginInput := entities.AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	dataUser, token, err := repositories.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error login, intrnal server error"))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login success", map[string]any{
		"token": token,
		"email": dataUser.Email,
		"id":    dataUser.Id,
	}))
}
