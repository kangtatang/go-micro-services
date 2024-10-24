package main

import (
	"employee-service/database"
	"employee-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi database
	database.Connect()

	// Set up routes
	routes.SetupRoutes(app)

	// Start server
	app.Listen(":3000")
}
