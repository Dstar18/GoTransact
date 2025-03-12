package main

import (
	"GoTransact/config"
	"GoTransact/controllers"
	"GoTransact/middleware"
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

	// Middleware JWT
	protectedJWT := e.Group("/api")
	protectedJWT.Use(middleware.JWTMiddleware)

	// public route
	e.POST("/login", controllers.Login)

	// protected route
	protectedJWT.GET("/customers", controllers.GetCustomers)

	e.Logger.Fatal(e.Start(":3000"))
}
