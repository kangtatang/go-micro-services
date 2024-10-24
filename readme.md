# Description

This is a simple backend to show the micro services architecture.

## Technologies

- Golang
- Go Fiber
- SQLite

## Using poly repo

- Every Service runs in a different port (go-ms-karyawan:3000, go-ms-project:4000, go-ms-mainapp:5000)
- Every Service has its own database (sqlite)
- Every Service has its own folder
- Use Post man to test the API
- DB Auto Create

## Installing
- run go mod tidy in every folder
- run go run main.go in every folder

## API Documentation
### go-ms-karyawan
- http://localhost:3000/api/employees (GET ALL)
- http://localhost:3000/api/employees (POST)
- http://localhost:3000/api/employees/:id (GET BY ID)
- http://localhost:3000/api/employees/:id (PUT)
- http://localhost:3000/api/employees/:id (DELETE)

Post request body example:

```json
{
    "name": "John Doe",
    "age": 30,
    "position": "Software Engineer"
}
```

### go-ms-project

- http://localhost:4000/api/projects (GET ALL)
- http://localhost:4000/api/projects (POST)
- http://localhost:4000/api/projects/:id (GET BY ID)
- http://localhost:4000/api/projects/:id (PUT)
- http://localhost:4000/api/projects/:id (DELETE)

Post request body example:

```json
{
    "name": "New Mobile Apps Project",
    "description": "Building a new mobile App for client",
    "employee_id": 1
}
```

### go-ms-mainapp

- http://localhost:5000/api/employee-projects/:id_employee (GET ALL Project By Employee ID)


Using this Concept, you can scale the micro service architecture.

find hard to set your own postman collection? 
Send me an email to: *kang.tatang@yahoo.co.id*
I will send you my postman collection file.
