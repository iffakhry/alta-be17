package handler

import (
	"be17/cleanarch/features/user"
	"be17/cleanarch/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	// memanggil func di repositories
	results, err := handler.userService.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var usersResponse []UserResponse
	for _, value := range results {
		usersResponse = append(usersResponse, UserResponse{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", usersResponse))

}
