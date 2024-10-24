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
curl -X POST http://localhost:3000/api/employees \
-H "Content-Type: application/json" \
-d '{
    "name": "John Doe",
    "age": 30,
    "position": "Software Engineer"
}'

```

## END POINT

Base url: http://localhost:3000

- Tambahkan karyawan: `POST /api/employees`
- Lihat semua karyawan: `GET /api/employees`
- Lihat karyawan berdasarkan ID: `GET /api/employees/:id`
- Update karyawan: `PUT /api/employees/:id`
- Hapus karyawan: `DELETE /api/employees/:id`

