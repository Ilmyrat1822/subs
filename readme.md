# Subscription Management Service

A robust backend service built with Go for managing user subscriptions and calculating total costs over selected periods.

## Features

- **Complete CRUD Operations** for subscription management
- **Cost Calculation** for selected time periods
- **RESTful API** with comprehensive endpoints
- **Swagger Documentation** for easy API exploration
- **PostgreSQL** database integration
- **Docker Support** for containerized deployment

## Tech Stack

- **Framework**: Echo (Go web framework)
- **ORM**: GORM
- **Database**: PostgreSQL
- **Documentation**: Swagger/OpenAPI
- **Containerization**: Docker & Docker Compose

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/subs` | Create a new subscription |
| `GET` | `/api/subs/{id}` | Get subscription by ID |
| `GET` | `/api/subs/list` | List all subscriptions |
| `PUT` | `/api/subs/{id}` | Update an existing subscription |
| `DELETE` | `/api/subs/{id}` | Delete a subscription |
| `GET` | `/api/subs/total` | Calculate total cost for a period |

## Getting Started

### Prerequisites

- Go 1.24
- PostgreSQL 14 or higher
- Docker and Docker Compose (optional)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd subscription-management-service
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your database credentials and configuration.

3. **Choose your setup method:**

   **Option A: Using Docker (Recommended)**
   ```bash
   docker-compose up --build
   ```

   **Option B: Manual Setup**
   
   Start PostgreSQL, then run:
   ```bash
   go mod download
   go run main.go
   ```

### Configuration

Edit your `.env` file with the following variables:

```env
POSTGRES_URI=postgres://postgres:password@localhost:5432/subsdb?sslmode=disable
DISABLE_AUTO_MIGRATION=false
PORT=7777
```

## API Documentation

Once the service is running, access the interactive Swagger UI at:

```
http://localhost:7777/swagger/index.html
```

### Generating Swagger Documentation

If you make changes to API endpoints, regenerate the docs:

```bash
swag init
```
## Docker Commands

```bash
# Start services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# Rebuild after changes
docker-compose up --build
```