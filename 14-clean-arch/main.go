package main

import (
	"be17/cleanarch/app/config"
	"be17/cleanarch/app/database"
	_userData "be17/cleanarch/features/user/data"
	_userHandler "be17/cleanarch/features/user/handler"
	_userService "be17/cleanarch/features/user/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	// dbPosgres := database.InitDBPosgres(cfg)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	userData := _userData.New(dbMysql)
	// userData := _userData.NewRaw(dbMysql)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.GET("/users", userHandlerAPI.GetAllUser)

	e.Logger.Fatal(e.Start(":8080"))
}
