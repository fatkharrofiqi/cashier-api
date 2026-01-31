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
│   ├── config/           # Configuration and database connection
│   ├── dto/              # Data transfer objects (request/response)
│   ├── handler/          # HTTP handlers for API endpoints
│   ├── model/            # Data models (database entities)
│   ├── repository/       # Data access layer (PostgreSQL)
│   ├── route/            # HTTP route definitions
│   └── service/          # Business logic layer
├── migrations/           # Database migration files
├── Makefile            # Build and migration commands
├── .env                # Environment variables
├── go.mod
└── .gitignore
```

## Features

- Product management with full CRUD operations
- Category management with full CRUD operations
- Products can be assigned to categories
- Clean architecture pattern (Handler → Service → Repository)
- PostgreSQL database with migration support
- DTO separation for clean API layer
- Graceful server shutdown

## Getting Started

### Prerequisites

- Go 1.24.1 or higher
- PostgreSQL database
- golang-migrate CLI tool

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd cashier-api
```

2. Install golang-migrate:
```bash
make install-migrate
```

3. Configure environment:
Create a `.env` file with your database connection:
```bash
DB_CONN=postgresql://user:password@localhost:5432/dbname?sslmode=disable
PORT=8080
```

4. Run migrations:
```bash
make migrateup
```

5. Build and run the application:
```bash
make start
```

Or use the Makefile:
```bash
make help          # Show all available commands
make build         # Build the application
make start         # Build and start the application
make migrateup     # Run database migrations up
make migratedown   # Rollback last migration
```

## API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/product` | Get all products |
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
  "stock": 10,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Makanan",
    "description": "Kategori makanan ringan"
  }
}
```

### Categories

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/category` | Get all categories |
| POST | `/api/category` | Create a new category |
| GET | `/api/category/{id}` | Get category by ID |
| PUT | `/api/category/{id}` | Update category by ID |
| DELETE | `/api/category/{id}` | Delete category by ID |

### Category Model

```json
{
  "id": 1,
  "name": "Makanan",
  "description": "Kategori makanan ringan"
}
```

### Example Usage

#### Get all products
```bash
curl http://localhost:8080/api/product
```

#### Create a product with category
```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Coca Cola",
    "price": 5000,
    "stock": 20,
    "category_id": 2
  }'
```

#### Create a category
```bash
curl -X POST http://localhost:8080/api/category \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Minuman",
    "description": "Kategori minuman"
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

- **Handler**: Handles HTTP requests and responses using DTOs
- **Service**: Contains business logic
- **Repository**: Manages data access using PostgreSQL
- **Model**: Defines database entities
- **DTO**: Defines API request/response structures

### Data Flow

```
HTTP Request → Handler (DTOs) → Service (Models) → Repository (DB) → PostgreSQL
                  ↓
           Response DTOs
```

## Migrations

The project uses golang-migrate for database schema management. Migration files are located in the `migrations/` directory.

To create a new migration:
```bash
make migrate-create NAME=add_new_table
```

To view migration status:
```bash
make migrate-force VERSION=1  # Force to specific version
```

## License

MIT
