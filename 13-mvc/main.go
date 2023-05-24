package main

import (
	"be17/mvc/config"
	"be17/mvc/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	// start server di port
	e.Logger.Fatal(e.Start(":8080"))
}
