package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-management-service/database"
	"project-management-service/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

// GetProjects godoc
// @Summary Get list of projects
// @Description Get all projects or projects based on employee ID
// @Tags projects
// @Param employee_id query int false "Employee ID"
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} models.ErrorResponse
// @Router /projects [get]
func GetProjects(c *fiber.Ctx) error {
	employeeID := c.Query("employee_id") // Dapatkan employee_id dari query parameter, jika ada

	var rows *sql.Rows
	var err error

	// Jika employee_id disertakan dalam query
	if employeeID != "" {
		rows, err = database.DB.Query("SELECT id, name, description, employee_id FROM projects WHERE employee_id = ?", employeeID)
	} else {
		// Jika tidak ada employee_id, ambil semua proyek
		rows, err = database.DB.Query("SELECT id, name, description, employee_id FROM projects")
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve projects"})
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.EmployeeID); err != nil {
			return err
		}
		projects = append(projects, project)
	}
	return c.JSON(projects)
}

// GetProjectByID godoc
// @Summary Get project by ID
// @Description Get a single project by ID along with employee details
// @Tags projects
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} models.CreateProjectResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id} [get]
func GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var project models.Project

	// Ambil data proyek dari database
	err := database.DB.QueryRow("SELECT id, name, description, employee_id FROM projects WHERE id = ?", id).Scan(&project.ID, &project.Name, &project.Description, &project.EmployeeID)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve project"})
	}

	// Ambil detail karyawan berdasarkan employee_id
	employee, err := fetchEmployee(project.EmployeeID)
	if err != nil || employee.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}

	// Menyusun response dengan detail proyek dan karyawan
	response := fiber.Map{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
		"employee":    employee, // Menambahkan detail karyawan di sini
	}

	return c.JSON(response)
}

// UpdateProject godoc
// @Summary Update project by ID
// @Description Update a project by ID
// @Tags projects
// @Param id path int true "Project ID"
// @Param project body models.Project true "Project Data"
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id} [put]
func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	_, err := database.DB.Exec("UPDATE projects SET name = ?, description = ?, employee_id = ? WHERE id = ?", project.Name, project.Description, project.EmployeeID, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update project"})
	}

	return c.JSON(fiber.Map{"message": "Project updated successfully"})
}

// DeleteProject godoc
// @Summary Delete project by ID
// @Description Delete a project by ID
// @Tags projects
// @Param id path int true "Project ID"
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects/{id} [delete]
func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete project"})
	}

	return c.JSON(fiber.Map{"message": "Project deleted successfully"})
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project and link it to an employee
// @Tags projects
// @Param project body models.Project true "Project Data"
// @Produce json
// @Success 201 {object} models.CreateProjectResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /projects [post]

func CreateProject(c *fiber.Ctx) error {
	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Ambil data karyawan dari service karyawan
	employee, err := fetchEmployee(project.EmployeeID)
	if err != nil || employee.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	// Insert project ke database
	_, err = database.DB.Exec("INSERT INTO projects (name, description, employee_id) VALUES (?, ?, ?)",
		project.Name, project.Description, project.EmployeeID)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(fiber.Map{
		"message":  "project created successfully",
		"project":  project,
		"employee": employee,
	})
}

// fetchEmployee retrieves employee details from the employee-service by the provided employeeID.
// Returns an Employee struct and any error encountered during the request.
func fetchEmployee(employeeID int) (*Employee, error) {
	// Konversi employeeID ke string
	resp, err := http.Get("http://localhost:3000/api/employees/" + strconv.Itoa(employeeID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch employee with status: %s", resp.Status)
	}

	employee := new(Employee)
	if err := json.NewDecoder(resp.Body).Decode(employee); err != nil {
		return nil, err
	}

	return employee, nil
}
