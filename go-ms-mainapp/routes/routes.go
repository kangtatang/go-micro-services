package routes

import (
	"employee-project-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Endpoint untuk mendapatkan employee dan project terkait berdasarkan employee_id
	app.Get("/api/employee-projects/:id", handlers.GetEmployeeProjects)
}
