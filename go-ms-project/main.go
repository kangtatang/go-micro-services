// package main

// import (
// 	"log"
// 	"project-management-service/database"
// 	"project-management-service/handlers"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	// Koneksi ke database
// 	database.Connect()

// 	// Route untuk membuat proyek baru
// 	app.Post("/api/projects", handlers.CreateProject)

// 	log.Fatal(app.Listen(":4000")) // Service ini berjalan di port 4000
// }

package main

import (
	"log"
	"project-management-service/database"
	"project-management-service/routes"

	// "github.com/gofiber/fiber/v2"
	_ "project-management-service/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Project Management Service API
// @version 1.0
// @description API for managing employees
// @host localhost:4000
// @BasePath /api
func main() {
	app := fiber.New()

	// Middleware Logger
	app.Use(logger.New())

	// Koneksi ke database
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Menjalankan server di port 4000
	log.Fatal(app.Listen(":4000")) // Service ini berjalan di port 4000
}
