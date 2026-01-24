# Cashier API

Simple cashier API built with Go using clean architecture pattern.

## Project Structure

```
cashier-api/
├── cmd/
│   └── http/
│       └── main.go       # Application entry point
├── internal/
│   ├── app/              # Application initialization and configuration
│   ├── handler/          # HTTP handlers for API endpoints
│   ├── model/            # Data models
│   ├── repository/       # Data access layer (in-memory storage)
│   ├── route/            # HTTP route definitions
│   └── service/          # Business logic layer
├── go.mod
└── .gitignore
```

## Features

- Product management with full CRUD operations
- Clean architecture pattern (Handler → Service → Repository)
- In-memory data storage
- Graceful server shutdown

## Getting Started

### Prerequisites

- Go 1.24.1 or higher

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd cashier-api
```

2. Run the application:
```bash
go run cmd/http/main.go
```

The server will start on port `8080` by default. You can change the port by setting the `PORT` environment variable:

```bash
PORT=3000 go run cmd/http/main.go
```

## API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/product` | Get all product |
| POST | `/api/product` | Create a new product |
| GET | `/api/product/{id}` | Get product by ID |
| PUT | `/api/product/{id}` | Update product by ID |
| DELETE | `/api/product/{id}` | Delete product by ID |

### Product Model

```json
{
  "id": 1,
  "name": "Indomie Goreng",
  "price": 3500,
  "stock": 10
}
```

### Example Usage

#### Get all product
```bash
curl http://localhost:8080/api/product
```

#### Create a product
```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Coca Cola",
    "price": 5000,
    "stock": 20
  }'
```

#### Get product by ID
```bash
curl http://localhost:8080/api/product/1
```

#### Update product
```bash
curl -X PUT http://localhost:8080/api/product/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Indomie Goreng Special",
    "price": 4000,
    "stock": 15
  }'
```

#### Delete product
```bash
curl -X DELETE http://localhost:8080/api/product/1
```

## Architecture

This project follows the clean architecture pattern:

- **Handler**: Handles HTTP requests and responses
- **Service**: Contains business logic
- **Repository**: Manages data access (currently using in-memory storage)
- **Model**: Defines data structures

## License

MIT
