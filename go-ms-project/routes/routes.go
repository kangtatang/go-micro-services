package routes

import (
	"project-management-service/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mendefinisikan semua route yang ada di Project Management Service
func SetupRoutes(app *fiber.App) {
	// Route untuk membuat proyek baru
	app.Post("/api/projects", handlers.CreateProject)

	// Route untuk mendapatkan semua proyek
	app.Get("/api/projects", handlers.GetProjects)

	// Route untuk mendapatkan proyek berdasarkan ID
	app.Get("/api/projects/:id", handlers.GetProjectByID)

	// Route untuk memperbarui proyek berdasarkan ID
	app.Put("/api/projects/:id", handlers.UpdateProject)

	// Route untuk menghapus proyek berdasarkan ID
	app.Delete("/api/projects/:id", handlers.DeleteProject)
}
