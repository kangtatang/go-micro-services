package main

import (
	"employee-project-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server di port 5000
	app.Listen(":5000")
}
