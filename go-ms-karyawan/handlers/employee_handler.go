package handlers

import (
	"employee-service/models"
	"employee-service/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	err := repository.CreateEmployee(*employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create employee"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "employee created successfully"})
}

func GetEmployees(c *fiber.Ctx) error {
	employees, err := repository.GetEmployees()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch employees"})
	}
	return c.JSON(employees)
}

func GetEmployeeByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee ID"})
	}

	employee, err := repository.GetEmployeeByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "employee not found"})
	}
	return c.JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee ID"})
	}

	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	err = repository.UpdateEmployee(id, *employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update employee"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "employee updated successfully"})
}

func DeleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid employee ID"})
	}

	err = repository.DeleteEmployee(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete employee"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "employee deleted successfully"})
}
