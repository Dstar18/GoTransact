package main

import (
	"GoTransact/config"
	"GoTransact/controllers"
	"GoTransact/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize logger
	utils.InitLogger()

	// connect database
	config.InitDB()
	// Migrate
	config.Migrate()
	// Seeder
	config.Seed()

	// initialize echo
	e := echo.New()

	e.POST("/login", controllers.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
