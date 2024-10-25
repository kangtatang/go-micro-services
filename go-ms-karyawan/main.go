package main

import (
	"employee-service/database"
	"employee-service/routes"

	_ "employee-service/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Employee Service API
// @version 1.0
// @description API for managing employees
// @host localhost:3000
// @BasePath /api
func main() {
	app := fiber.New()

	// Middleware Logger
	app.Use(logger.New())

	// Inisialisasi database
	database.Connect()

	// Set up routes
	routes.SetupRoutes(app)

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Start server
	app.Listen(":3000")
}
