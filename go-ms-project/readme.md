## Install dependency

```bash
   go get github.com/gofiber/fiber/v2
   go get modernc.org/sqlite
```
Or just Type and hit:

```bash
go mod tidy
```

## curl post

```bash
curl -X POST http://localhost:4000/api/projects \
-H "Content-Type: application/json" \
-d '{
    "name": "New Mobile Apps Project",
    "description": "Building a new mobile App for client",
    "employee_id": 1
}'

```

## END POINT

Base url: http://localhost:4000

- Add Project: `POST /api/projects`
- Lihat semua project: `GET /api/projects`
- Lihat project berdasarkan ID: `GET /api/projects/:id`
- Update project: `PUT /api/projects/:id`
- Hapus project: `DELETE /api/projects/:id`

