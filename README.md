# Classroom Management System

A simple REST API for managing classrooms and student enrollment, written in Go with PostgreSQL.

## Requirements

- Go 1.27+
- PostgreSQL

## Setup

1. Create a PostgreSQL database and the required tables:

   ```sql
   CREATE TABLE classrooms (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       capacity INT NOT NULL
   );

   CREATE TABLE students (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       email TEXT NOT NULL,
       classroom_id INT REFERENCES classrooms(id)
   );
   ```

2. Set the database connection string (currently hardcoded in [main.go](cmd/api/main.go) — update it to match your local PostgreSQL setup):

   ```
   postgres://user:password@localhost/classroom_db?sslmode=disable
   ```

3. Install dependencies and run the API:

   ```bash
   go mod download
   go run ./cmd/api
   ```

The server starts on `http://localhost:8080`.

## API Endpoints

### Create a classroom

```
POST /api/classrooms
Content-Type: application/json

{
  "name": "Grade 5A",
  "capacity": 30
}
```

### Enroll a student

```
POST /api/students
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "classroom_id": 1
}
```

## Project Structure

```
cmd/api/            Application entry point
internal/database/  Database connection setup
internal/handlers/  HTTP request handlers
internal/models/    Data models and queries
```
