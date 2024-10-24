package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EmployeeID  int    `json:"employee_id"`
}

type EmployeeProjectsResponse struct {
	Employee Employee  `json:"employee"`
	Projects []Project `json:"projects"`
}

// GetEmployeeProjects - Handler untuk mendapatkan data karyawan dan proyek terkait
func GetEmployeeProjects(c *fiber.Ctx) error {
	employeeID := c.Params("id")

	// Ambil data karyawan dari employee-service
	employee, err := fetchEmployee(employeeID)
	if err != nil || employee.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}

	// Ambil data proyek dari project-management-service
	projects, err := fetchProjectsByEmployee(employeeID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve projects"})
	}

	// Gabungkan data karyawan dan proyek dalam response
	response := EmployeeProjectsResponse{
		Employee: *employee,
		Projects: projects,
	}

	return c.JSON(response)
}

// Fungsi untuk mengambil data karyawan dari employee service
func fetchEmployee(employeeID string) (*Employee, error) {
	resp, err := http.Get("http://localhost:3000/api/employees/" + employeeID) // Mengacu ke employee service
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	employee := new(Employee)
	if err := json.NewDecoder(resp.Body).Decode(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

// Fungsi untuk mengambil data proyek dari project-management-service
func fetchProjectsByEmployee(employeeID string) ([]Project, error) {
	resp, err := http.Get("http://localhost:4000/api/projects?employee_id=" + employeeID) // Mengacu ke project management service
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var projects []Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}
