package routes

import (
	"be17/mvc/controllers"
	"be17/mvc/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// define route / endpoint
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users", controllers.GetAllUserController, middlewares.JWTMiddleware())

	e.POST("/login", controllers.LoginController)

	return e
}
