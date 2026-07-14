# Student REST API with Go

A REST API for managing student records, built with Go's standard `net/http`
router and SQLite. The application supports creating, reading, updating, and
deleting students, validates incoming JSON, and loads its settings from YAML.

## Features

- Student CRUD endpoints
- SQLite persistence with automatic table creation
- Request validation using `go-playground/validator`
- YAML configuration with `cleanenv`
- Structured application logging
- Graceful HTTP server shutdown

## Requirements

- Go `1.26.3` or a compatible version
- A C compiler, such as GCC, for the `go-sqlite3` CGO dependency
- `curl` or another HTTP client for trying the API

## Getting Started

Clone the repository and enter the project directory:

```bash
git clone https://github.com/mdiktushar/REST-API-with-GoLang.git
cd REST-API-with-GoLang
```

Download the dependencies:

```bash
go mod download
```

Create a local configuration file and the SQLite storage directory:

```bash
cp config/example.yaml config/local.yaml
mkdir -p storage
```

The example configuration starts the server at `http://localhost:8082` and
stores data in `storage/storage.db`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

Run the API by passing the configuration file as a flag:

```bash
go run ./cmd/REST-API-with-GoLang -config config/local.yaml
```

You can alternatively set `CONFIG_PATH`:

```bash
CONFIG_PATH=config/local.yaml go run ./cmd/REST-API-with-GoLang
```

The `students` table is created automatically when the application starts.
Stop the server gracefully with `Ctrl+C`.

## API Endpoints

The base URL for the example configuration is `http://localhost:8082`.

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/api/students` | Create a student |
| `GET` | `/api/students` | Get all students |
| `GET` | `/api/students/{id}` | Get a student by ID |
| `PUT` | `/api/students/{id}` | Update a student |
| `DELETE` | `/api/students/{id}` | Delete a student |

### Create a student

```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Tushar","email":"tushar@gmail.com","age":15}'
```

Response (`201 Created`):

```json
{
  "id": 1
}
```

The `name`, `email`, and `age` fields are required.

### Get all students

```bash
curl http://localhost:8082/api/students
```

Response (`200 OK`):

```json
[
  {
    "id": 1,
    "name": "Tushar",
    "email": "tushar@gmail.com",
    "age": 15
  }
]
```

### Get a student by ID

```bash
curl http://localhost:8082/api/students/1
```

Response (`200 OK`):

```json
{
  "id": 1,
  "name": "Tushar",
  "email": "tushar@gmail.com",
  "age": 15
}
```

### Update a student

```bash
curl -X PUT http://localhost:8082/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Tushar Ahmed","email":"tushar@gmail.com","age":16}'
```

Response (`200 OK`):

```json
{
  "id": 1,
  "name": "Tushar Ahmed",
  "email": "tushar@gmail.com",
  "age": 16
}
```

### Delete a student

```bash
curl -X DELETE http://localhost:8082/api/students/1
```

Response (`200 OK`):

```json
{
  "message": "Student deleted successfully"
}
```

### Error response

Invalid requests return a JSON response in this format:

```json
{
  "status": "Error",
  "error": "field Name is required field"
}
```

## Project Structure

```text
.
├── cmd/REST-API-with-GoLang/
│   └── main.go                  # Application entry point and routes
├── config/
│   └── example.yaml            # Example application configuration
├── internal/
│   ├── config/                 # Configuration loading
│   ├── http/handlers/student/  # Student HTTP handlers
│   ├── storage/                # Storage interface
│   │   └── sqlite/             # SQLite implementation
│   ├── types/                  # Student model
│   └── utils/response/         # JSON response helpers
├── go.mod
└── README.md
```

## Development

Format the code:

```bash
go fmt ./...
```

Run all tests and compile every package:

```bash
go test ./...
```

Build the API:

```bash
go build ./cmd/REST-API-with-GoLang
```

## License

No license has been added yet.
