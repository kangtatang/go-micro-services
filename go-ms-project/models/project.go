package models

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EmployeeID  int    `json:"employee_id"` // ID karyawan yang diambil dari service karyawan
}
