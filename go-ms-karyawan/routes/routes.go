package routes

import (
	"employee-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Employee routes
	api.Post("/employees", handlers.CreateEmployee)
	api.Get("/employees", handlers.GetEmployees)
	api.Get("/employees/:id", handlers.GetEmployeeByID)
	api.Put("/employees/:id", handlers.UpdateEmployee)
	api.Delete("/employees/:id", handlers.DeleteEmployee)
}
