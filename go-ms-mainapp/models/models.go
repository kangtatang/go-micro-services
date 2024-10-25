// package models

// // ErrorResponse defines the structure for error responses
// type ErrorResponse struct {
// 	Error string `json:"error"`
// }

// // Employee defines the structure for an employee
// type Employee struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Position string `json:"position"`
// }

// // Project defines the structure for a project
// type Project struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	EmployeeID  int    `json:"employee_id"`
// }

// // EmployeeProjectsResponse defines the response structure for employee projects
// type EmployeeProjectsResponse struct {
// 	Employee Employee  `json:"employee"`
// 	Projects []Project `json:"projects"`
// }

// models/models.go

package models

// Employee struct untuk menyimpan data karyawan
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

// Project struct untuk menyimpan data proyek
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EmployeeID  int    `json:"employee_id"`
}

// EmployeeProjectsResponse struct untuk menyimpan response karyawan dan proyek
type EmployeeProjectsResponse struct {
	Employee Employee  `json:"employee"`
	Projects []Project `json:"projects"`
}

// ErrorResponse struct untuk menyimpan response error
type ErrorResponse struct {
	Error string `json:"error"`
}
