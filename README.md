# Go Fiber GORM Backend API with Container Pattern

A simple and efficient backend API built with Go, utilizing the Fiber web framework and GORM for database operations, implementing the container pattern for dependency injection and clean architecture.

## Features

- RESTful API endpoints
- Container pattern implementation for dependency injection
- Clean and maintainable code structure
- Fast HTTP routing with Fiber
- Database abstraction with GORM
- Modular architecture
- Easy configuration management
- Middleware support
- Error handling
- JSON request/response handling

## Tech Stack

- **Go** - Programming language
- **Fiber** - Express-inspired web framework
- **GORM** - Object-relational mapping library
- **Container Pattern** - Dependency injection pattern
- **Docker** - Used to run the database

## Installation

### Prerequisites

- Go 1.19 or higher
- Database - PostgreSQL

### Setup

1. Clone the repository
2. run "go mod tidy"
3. ensure you have all the .env variables set, according to the .env.example file provided
3. either compile the project and run it or "go run ./cmd/api"