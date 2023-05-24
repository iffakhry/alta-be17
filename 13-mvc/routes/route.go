package routes

import (
	"be17/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// define route / endpoint
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users", controllers.GetAllUserController)

	return e
}
