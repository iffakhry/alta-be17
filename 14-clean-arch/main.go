package main

import (
	"be17/cleanarch/app/config"
	"be17/cleanarch/app/database"
	"be17/cleanarch/app/router"

	// _bookData "be17/cleanarch/features/book/data"
	// _bookHandler "be17/cleanarch/features/book/handler"
	// _bookService "be17/cleanarch/features/book/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	// dbPosgres := database.InitDBPosgres(cfg)

	database.InitialMigration(dbMysql)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(dbMysql, e)

	// userData := _userData.New(dbMysql)
	// // userData := _userData.NewRaw(dbMysql)
	// userService := _userService.New(userData)
	// userHandlerAPI := _userHandler.New(userService)

	// e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	// e.POST("/users", userHandlerAPI.CreateUser)
	// e.POST("/login", userHandlerAPI.Login)

	e.Logger.Fatal(e.Start(":80"))
}
