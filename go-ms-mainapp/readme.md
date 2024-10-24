# Employee Project Service

## Description

This service is responsible for handling requests related to employees and their projects. It fetches employee data from the Employee Service and project data from the Project Management Service.

```bash
curl -X GET http://localhost:5000/api/employee-projects/1

```

## Contoh output

```json
{
    "employee": {
        "id": 1,
        "name": "John Doe",
        "age": 30,
        "position": "Software Engineer"
    },
    "projects": [
        {
            "id": 1,
            "name": "New Website Project",
            "description": "Building a new website for client",
            "employee_id": 1
        }
    ]
}

```