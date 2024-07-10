
# Golang CRUD API with PostgreSQL and JWT Authentication

This project is a demonstration of a CRUD API built using Golang, PostgreSQL for the database, and JWT for authentication. It follows best coding standards and aims to showcase professional Golang development skills.

## Project Structure

The project is structured as follows:

```plaintext
myapp/
├── cmd/
│   └── api/
│       ├── main.go
├── config/
│   └── config.go
├── internal/
│   ├── auth/
│   │   ├── jwt.go
│   │   └── middleware.go
│   ├── db/
│   │   ├── migration/
│   │   │   └── migrate.go
│   │   └── postgres.go
│   ├── handlers/
│   │   └── user.go
│   ├── models/
│   │   └── user.go
│   ├── repository/
│   │   └── user.go
│   ├── router/
│   │   └── router.go
│   └── utils/
│       ├── utils.go
│       └── logger.go
├── pkg/
│   └── logger/
│       └── logger.go
├── go.mod
└── go.sum
```
### Key Components

- **cmd/api/main.go**: Entry point of the application. Initializes configuration, database connection, and starts the server.
  
- **config/config.go**: Configuration loader using Viper, reads from `.env` file.

- **internal/auth/**: Contains JWT authentication logic and middleware.
  
- **internal/db/**: Database related logic including PostgreSQL connection and migration scripts.

- **internal/handlers/**: HTTP request handlers for different API endpoints.

- **internal/models/**: Data models used throughout the application.

- **internal/repository/**: Data access layer for interacting with the database.

- **internal/server/**: HTTP server setup using Chi router.

- **pkg/logger/**: Logging utility for the application.

## Setup and Installation

1. **Clone the repository**:
   `git clone <repository_url>`
   `cd myapp`

2. **Set up the database**:
   - Ensure PostgreSQL is installed and running.
   - Create a `.env` file in the root directory with the following configuration:
```plaintext     
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your_db_user
     DB_PASSWORD=your_db_password
     DB_NAME=your_db_name
     DB_SSLMODE=disable
     SERVER_PORT=8080
     JWT_SECRET_KEY=your_secret_key
```

3. **Database Migration**:
   - Run the database migration to set up the required tables:

     `go run cmd/api/main.go migrate`

4. **Start the API server**:
   - Start the Golang server which listens on `SERVER_PORT` defined in the `.env` file (default is 8080):

     `go run cmd/api/main.go`

## API Endpoints

### User Management

- **POST /register**: Register a new user.
  `{
  "username": "example",
  "password": "password",
  "first_name": "John",
  "last_name": "Doe",
  "email": "example@example.com",
  "phone_number": "+1234567890",
  "address": "123 Main St",
  "city": "Anytown",
  "state": "Anystate",
  "country": "Anycountry",
  "zip_code": "12345",
  }`

- **POST /login**: Authenticate and generate JWT token.
  `{
    "username": "example",
    "password": "password"
  }`
  Successful login response includes a JWT token in the `Authorization` header.

## Dependencies

- **github.com/go-chi/chi/v5**: Lightweight, idiomatic router for building Go HTTP services.
- **github.com/lib/pq**: PostgreSQL driver for Go.
- **github.com/spf13/viper**: Configuration management with support for `.env` files.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests for any improvements or features you'd like to see.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
