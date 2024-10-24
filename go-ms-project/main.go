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

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Koneksi ke database
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Menjalankan server di port 4000
	log.Fatal(app.Listen(":4000")) // Service ini berjalan di port 4000
}
