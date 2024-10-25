package handlers

import (
	"employee-service/models"
	"employee-service/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateEmployee godoc
// @Summary      Create a new employee
// @Description  Create a new employee with JSON body
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param        employee  body      models.Employee  true  "Employee data"
// @Success      201       {object}  map[string]string "message": "employee created successfully"
// @Failure      400       {object}  map[string]string "error": "cannot parse JSON"
// @Failure      500       {object}  map[string]string "error": "failed to create employee"
// @Router       /employees [post]

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

// GetEmployees godoc
// @Summary      Get all employees
// @Description  Retrieve a list of employees
// @Tags         Employees
// @Produce      json
// @Success      200  {array}   models.Employee
// @Failure      500  {object}  map[string]string "error": "failed to fetch employees"
// @Router       /employees [get]
func GetEmployees(c *fiber.Ctx) error {
	employees, err := repository.GetEmployees()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch employees"})
	}
	return c.JSON(employees)
}

// GetEmployeeByID godoc
// @Summary      Get employee by ID
// @Description  Retrieve an employee by their ID
// @Tags         Employees
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  models.Employee
// @Failure      400  {object}  map[string]string "error": "invalid employee ID"
// @Failure      404  {object}  map[string]string "error": "employee not found"
// @Router       /employees/{id} [get]
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

// UpdateEmployee godoc
// @Summary      Update an employee
// @Description  Update an existing employee by ID
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param        id        path      int             true  "Employee ID"
// @Param        employee  body      models.Employee true  "Updated Employee Data"
// @Success      200       {object}  map[string]string "message": "employee updated successfully"
// @Failure      400       {object}  map[string]string "error": "cannot parse JSON"
// @Failure      500       {object}  map[string]string "error": "failed to update employee"
// @Router       /employees/{id} [put]
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

// DeleteEmployee godoc
// @Summary      Delete an employee
// @Description  Delete an employee by ID
// @Tags         Employees
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  map[string]string "message": "employee deleted successfully"
// @Failure      400  {object}  map[string]string "error": "invalid employee ID"
// @Failure      500  {object}  map[string]string "error": "failed to delete employee"
// @Router       /employees/{id} [delete]
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
