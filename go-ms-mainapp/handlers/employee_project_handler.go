package handlers

import (
	"encoding/json"
	"net/http"

	"employee-project-service/models" // pastikan path ini sesuai dengan struktur modul Anda

	"github.com/gofiber/fiber/v2"
)

// GetEmployeeProjects - Handler untuk mendapatkan data karyawan dan proyek terkait
// @Summary Get employee projects
// @Description Get employee data and related projects
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.EmployeeProjectsResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /employee-projects/{id} [get]
func GetEmployeeProjects(c *fiber.Ctx) error {
	employeeID := c.Params("id")

	// Ambil data karyawan dari employee-service
	employee, err := fetchEmployee(employeeID)
	if err != nil || employee.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{Error: "Employee not found"})
	}

	// Ambil data proyek dari project-management-service
	projects, err := fetchProjectsByEmployee(employeeID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Could not retrieve projects"})
	}

	// Gabungkan data karyawan dan proyek dalam response
	response := models.EmployeeProjectsResponse{
		Employee: *employee,
		Projects: projects,
	}

	return c.JSON(response)
}

// Fungsi untuk mengambil data karyawan dari employee service
func fetchEmployee(employeeID string) (*models.Employee, error) {
	resp, err := http.Get("http://localhost:3000/api/employees/" + employeeID) // Mengacu ke employee service
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	employee := new(models.Employee) // Menggunakan models.Employee
	if err := json.NewDecoder(resp.Body).Decode(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

// Fungsi untuk mengambil data proyek dari project-management-service
func fetchProjectsByEmployee(employeeID string) ([]models.Project, error) {
	resp, err := http.Get("http://localhost:4000/api/projects?employee_id=" + employeeID) // Mengacu ke project management service
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var projects []models.Project // Menggunakan models.Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}
