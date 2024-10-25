package main

import (
	"employee-project-service/routes"

	// "github.com/gofiber/fiber/v2"
	_ "employee-project-service/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Employee Project Management Service API
// @version 1.0
// @description API for managing employees
// @host localhost:5000
// @BasePath /api
func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Middleware Logger
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Jalankan server di port 5000
	app.Listen(":5000")
}
