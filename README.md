# REST API with GoLang

A small Go project scaffold for building a student REST API.

At the moment, the application entry point prints a welcome message. The README is structured so the project can grow naturally as API routes, persistence, configuration, and tests are added.

## Project Structure

```text
.
├── cmd/
│   └── REST-API-with-GoLang/
│       └── main.go
├── go.mod
└── README.md
```

## Requirements

- Go `1.26.3` or compatible

## Getting Started

Clone the repository:

```bash
git clone https://github.com/mdiktushar/REST-API-with-GoLang.git
cd REST-API-with-GoLang
```

Run the application:

```bash
go run ./cmd/REST-API-with-GoLang
```

Expected output:

```text
Welcome to students api
```

## Development

Format the code:

```bash
go fmt ./...
```

Run tests:

```bash
go test ./...
```

Build the application:

```bash
go build ./cmd/REST-API-with-GoLang
```

## Planned API

This project is intended to become a REST API for managing student data. Possible future endpoints include:

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/students` | List all students |
| `GET` | `/students/{id}` | Get a single student |
| `POST` | `/students` | Create a student |
| `PUT` | `/students/{id}` | Update a student |
| `DELETE` | `/students/{id}` | Delete a student |

## Roadmap

- Add HTTP server setup
- Add student routes and handlers
- Add request and response models
- Add validation
- Add database integration
- Add unit and integration tests

## License

Add a license before publishing or distributing this project.
