# LineBK Backend Assignment

---

## 📋 Contents

- [📖 About](#-about)
- [✨ Features](#-features)
- [⚙️ Prerequisites](#️-prerequisites)
- [🚀 Quick Start](#-quick-start)
- [⚙️ Configuration](#️-configuration)
- [📚 API Documentation](#-api-documentation)
- [🗄️ Database](#️-database)
- [🧪 Testing](#-testing)
- [⚡ Performance Testing](#-performance-testing)
- [📁 Project Structure](#-project-structure)

---

## 📖 About

This backend API assignment focuses on building robust functionalities with a strong emphasis on comprehensive testing and a well-defined project structure.

**Technology Stack:**

- **Backend:** Go 1.24.4 with Gin framework
- **Database:** MySQL 8.0 with SQLx
- **Authentication:** JWT tokens
- **Documentation:** Swagger
- **Testing:** Unit tests with mocks, K6 performance testing
- **Deployment:** Docker & Docker Compose

---

## ✨ Features

- 🔐 **JWT Authentication** - Secure user authentication with access and refresh tokens
- 👤 **User Management** - User registration, profile management
- 📊 **Comprehensive Logging** - Structured logging with Zerolog
- 🧪 **Unit Testing** - Comprehensive test coverage with mocks
- 📈 **Performance Testing** - K6 load testing capabilities
- 📚 **API Documentation** - Auto-generated Swagger documentation
- 🐳 **Docker Support** - Containerized deployment with Docker Compose

---

## ⚙️ Prerequisites

### For Docker Deployment (Recommended)

- **Docker & Docker Compose** - [Install Docker](https://docs.docker.com/get-docker/)

### For Local Development (Optional)

If you want to run the API server locally (database can be Docker or local):

**Required:**

- **Go 1.24.4+** - [Download Go](https://golang.org/dl/)
- **Make** - Usually pre-installed on Unix systems
- **golang-migrate** - For database migrations
  ```bash
  go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```
- **swag** - For Swagger documentation generation
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

**Database Options:**

**Option A: Use Docker for MySQL (Recommended)**

- **Docker** - [Install Docker](https://docs.docker.com/get-docker/)
- Uses `make mysql` command (runs MySQL in Docker container)

**Option B: Local MySQL Installation**

- **MySQL 8.0+** - [Download MySQL](https://dev.mysql.com/downloads/)
- Update `DATABASE_URL` in your environment to point to local MySQL

**Optional:**

- **k6** - For performance testing - [Install k6](https://k6.io/docs/get-started/installation/)

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/WuttinunSkywalker/linebk-backend-assignment.git
cd linebk-backend-assignment
```

### 2. Run with Docker (Recommended)

**Start the application:**

```bash
docker-compose up -d
```

**Stop the application:**

```bash
docker-compose down
```

**Clean stop (remove volumes):**

```bash
docker-compose down -v
```

### 3. Access the Application

Once the application is running:

| Service            | Port   | Purpose              | Access                     |
| ------------------ | ------ | -------------------- | -------------------------- |
| **API Server**     | `8080` | Main application API | `http://localhost:8080`    |
| **MySQL Database** | `3306` | Database connection  | For local development only |

> 📋 See the [API Documentation](#-api-documentation) section for test credentials and detailed testing instructions.

### 4. Local Development Setup (Optional)

If you prefer to run the API server locally:

**Install dependencies:**

```bash
go mod download
```

**Option A: With Docker MySQL (Recommended)**

```bash
make mysql          # Start MySQL in Docker container
make migrateup      # Run migrations
make server         # Start API server
```

**Option B: With Local MySQL**

```bash
# 1. Start your local MySQL server
# 2. Create database 'assignment'
# 3. Update DATABASE_URL in app.env and Makefile if needed
make migrateup      # Run migrations
make server         # Start API server
```

**Useful commands:**

```bash
make swagger       # Regenerate Swagger docs
make test          # Run unit tests
make stress-test   # Run K6 performance tests
```

---

## ⚙️ Configuration

The application uses environment variables for configuration. Default values are provided in `app.env`:

### Environment Variables

| Variable                     | Description                              | Example                                                     |
| ---------------------------- | ---------------------------------------- | ----------------------------------------------------------- |
| `PORT`                       | Server port (change if 8080 is in use)   | `8080`                                                      |
| `LOG_LEVEL`                  | Logging level (debug, info, warn, error) | `debug`                                                     |
| `LOG_FORMAT`                 | Log format (console, json)               | `console`                                                   |
| `DATABASE_URL`               | MySQL connection string                  | `root:secret@tcp(127.0.0.1:3306)/assignment?parseTime=true` |
| `JWT_SECRET`                 | JWT signing secret                       | `secret`                                                    |
| `JWT_ISSUER`                 | JWT token issuer                         | `linebk-backend-assignment`                                 |
| `JWT_ACCESS_EXPIRY_SECONDS`  | Access token expiry in seconds           | `43200` (12 hours)                                          |
| `JWT_REFRESH_EXPIRY_SECONDS` | Refresh token expiry in seconds          | `86400` (24 hours)                                          |

---

## 📚 API Documentation

### Swagger UI

Once the server is running, access the interactive API documentation at:

```
http://localhost:8080/swagger/index.html
```

### Login Credentials

The application automatically seeds test data during startup (via [initial seed migration](migrations/000002_seed_init.up.sql)). Use these credentials for testing:

- **User ID:** `0befecd8-fccb-417e-aa0a-1a23c021f413`
- **PIN:** `123456`

> **Note:** All seeded users use the same PIN (`123456`). You can use any User ID from the seed data with PIN `123456` for testing.

### How to Test the API

1. **Login:** Use the test credentials above with `/auth/login` endpoint
2. **Get Token:** Copy the JWT access token from the login response
3. **Authenticate:** Click "Authorize" in Swagger UI and enter: `Bearer <your-jwt-token>`
4. **Test Endpoints:** All authenticated endpoints will now work in Swagger UI

> **Note:** The database is automatically populated with mock data for testing purposes when migrations run.

---

## 🗄️ Database

### ER Diagram

![image (1)](https://github.com/user-attachments/assets/c43d3a19-f49a-4b6c-a87f-184b668ace9f)

> **Note:** When using Docker Compose, migrations are automatically executed during container startup.

### Adding More Test Data

If you need more data than the default seed provides, you can download additional sample data and import it:

**Download Additional Data:**

- [Download Mock Data](https://drive.google.com/file/d/1F04D-DFyBDMQu8qHzkglehsjTnHe6PjS/view?usp=drive_link)

**Import via Docker:**

```bash
docker exec -i assignment-mysql mysql -u root -psecret assignment < mock_data.sql
```

**Import via Local MySQL:**

```bash
# If running MySQL locally
mysql -u root -p assignment < mock_data.sql
```

### Local Development Migrations

```bash
# Run all migrations
make migrateup

# Rollback migrations
make migratedown

# Use go-migrate directly
migrate -path migrations -database "mysql://root:secret@tcp(localhost:3306)/assignment" up
```

> **Note:** If using `make migrateup/migratedown` with a custom database connection, update the `DATABASE_URL` variable in the Makefile to match your database configuration.

---

## 🧪 Testing

### Unit Tests

**Option 1: Using Make (Recommended)**

```bash
# Generate mocks (if needed)
make mock

# Run all tests
make test
```

**Option 2: Manual Commands**

```bash
# Generate mocks (if needed)
mockery --all

# Run tests manually
go test ./...

# Run specific package tests
go test ./internal/api/auth/...
```

**Testing Stack:**

- **[Mockery](https://github.com/vektra/mockery)** - Automatic mock generation
- **[Testify](https://github.com/stretchr/testify)** - Assertions and test suites

**Mock Generation:**
The project uses Mockery to automatically generate mocks for all interfaces. Mocks are generated in `mocks/` directories alongside their respective interfaces.

---

## ⚡ Performance Testing

The project includes K6 stress testing:

```bash
# Run stress tests
make stress-test

# Run custom K6 script
k6 run k6/stress-test.js
```

### Stress Test Results

Here are the performance test results captured during K6 testing:

![image (3)](https://github.com/user-attachments/assets/14314e25-ed35-4081-a9b4-dbee588214a8)

_The above screenshot shows the K6 performance metrics including response times, throughput, and error rates during load testing._

---

## 📁 Project Structure

```
linebk-backend-assignment/
├── cmd/
│   └── main.go                   # Application entry point
├── internal/
│   ├── api/
│   │   ├── account/              # Account module
│   │   │   ├── dto.go            # Request/response DTOs
│   │   │   ├── handler.go        # HTTP handlers
│   │   │   ├── handler_test.go   # Handler tests
│   │   │   ├── model.go          # Domain models
│   │   │   ├── repository.go     # Data access
│   │   │   ├── usecase.go        # Business logic
│   │   │   ├── usecase_test.go   # Usecase tests
│   │   │   └── mocks/            # Generated mocks
│   │   ├── auth/                 # Authentication module
│   │   ├── banner/               # Banner module
│   │   ├── debit/                # Debit module
│   │   ├── transaction/          # Transaction module
│   │   └── user/                 # User module
│   ├── config/
│   │   └── config.go             # Configuration management
│   ├── database/
│   │   └── database.go           # Database connection
│   ├── middleware/
│   │   ├── auth.go               # Authentication middleware
│   │   └── error.go              # Error handling middleware
│   └── routes/
│       └── routes.go             # Route definitions
├── pkg/
│   ├── errs/                     # Custom api errors
│   ├── logger/                   # Logging utilities
│   ├── pagination/               # Pagination utilities
│   ├── response/                 # Response formatting
│   └── token/                    # JWT token utilities
├── migrations/                   # Database migrations
├── docs/                         # Generated Swagger docs
├── k6/
│   └── stress-test.js            # K6 stress test
├── docker-compose.yml            # Docker services
├── Dockerfile                    # Container definition
├── Makefile                      # Commands
├── app.env                       # Environment variables
├── go.mod                        # Go dependencies
└── README.md                     # Project README file
```

---
