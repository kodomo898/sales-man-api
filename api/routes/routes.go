package routes

import (
    healthcontrollers "jellness-api/api/controllers/health"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init is Routing
func Init() {

	e := echo.New()
	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", healthcontrollers.Health())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
