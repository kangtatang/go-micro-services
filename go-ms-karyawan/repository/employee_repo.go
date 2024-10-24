package repository

import (
	"employee-service/database"
	"employee-service/models"
	"log"
)

func CreateEmployee(employee models.Employee) error {
	query := "INSERT INTO employees (name, age, position) VALUES (?, ?, ?)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(employee.Name, employee.Age, employee.Position)
	return err
}

func GetEmployees() ([]models.Employee, error) {
	rows, err := database.DB.Query("SELECT id, name, age, position FROM employees")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	employees := []models.Employee{}

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Position)
		if err != nil {
			log.Println(err)
			continue
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func GetEmployeeByID(id int) (models.Employee, error) {
	var employee models.Employee
	query := "SELECT id, name, age, position FROM employees WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Position)
	return employee, err
}

func UpdateEmployee(id int, employee models.Employee) error {
	query := "UPDATE employees SET name = ?, age = ?, position = ? WHERE id = ?"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(employee.Name, employee.Age, employee.Position, id)
	return err
}

func DeleteEmployee(id int) error {
	query := "DELETE FROM employees WHERE id = ?"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
