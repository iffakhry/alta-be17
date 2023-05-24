package main

import (
	"be17/mvc/config"
	"be17/mvc/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	// middlewares.LogMiddlewares(e)
	// start server di port
	e.Logger.Fatal(e.Start(":8080"))
}
