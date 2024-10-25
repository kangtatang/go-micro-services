package models

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

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type EmployeeProjectsResponse struct {
	Employee Employee  `json:"employee"`
	Projects []Project `json:"projects"`
}
