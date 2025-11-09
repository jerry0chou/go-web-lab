# Gin GORM Project

A simple Go web application using Gin framework, GORM ORM, and PostgreSQL database.

## Features

- Student, Teacher, Course, and Enrollment models with relationships
- RESTful API endpoints for CRUD operations
- PostgreSQL database with GORM migrations

## Docker Setup

### Prerequisites

- Docker
- Docker Compose

### Running with Docker Compose

1. Build and start all services:
```bash
docker-compose up -d
```

2. View logs:
```bash
docker-compose logs -f app
```

3. Stop services:
```bash
docker-compose down
```

4. Stop and remove volumes (clears database data):
```bash
docker-compose down -v
```

### Building Docker Image Manually

```bash
docker build -f gin_gorm/Dockerfile -t gin-gorm-app ..
```

### Running Container Manually

First, start PostgreSQL:
```bash
docker run -d \
  --name postgres_db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=school_db \
  -p 5432:5432 \
  postgres:16-alpine
```

Then run the application:
```bash
docker run -d \
  --name gin_gorm_app \
  -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  -e DB_NAME=school_db \
  -e DB_PORT=5432 \
  gin-gorm-app
```

## API Endpoints

### Students
- `GET /students` - Get all students
- `GET /students/:id` - Get student by ID
- `POST /students` - Create student
- `PUT /students/:id` - Update student
- `DELETE /students/:id` - Delete student

### Teachers
- `GET /teachers` - Get all teachers
- `GET /teachers/:id` - Get teacher by ID
- `POST /teachers` - Create teacher
- `PUT /teachers/:id` - Update teacher
- `DELETE /teachers/:id` - Delete teacher

### Courses
- `GET /courses` - Get all courses
- `GET /courses/:id` - Get course by ID
- `POST /courses` - Create course
- `PUT /courses/:id` - Update course
- `DELETE /courses/:id` - Delete course

### Enrollments
- `GET /enrollments` - Get all enrollments
- `GET /enrollments/:id` - Get enrollment by ID
- `POST /enrollments` - Create enrollment
- `PUT /enrollments/:id` - Update enrollment
- `DELETE /enrollments/:id` - Delete enrollment

## Environment Variables

- `DB_HOST` - Database host (default: localhost)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: school_db)
- `DB_PORT` - Database port (default: 5432)

## Example API Calls

Create a student:
```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":20}'
```

Create a teacher:
```bash
curl -X POST http://localhost:8080/teachers \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Smith","email":"jane@example.com","subject":"Mathematics"}'
```

Create a course:
```bash
curl -X POST http://localhost:8080/courses \
  -H "Content-Type: application/json" \
  -d '{"title":"Calculus 101","description":"Introduction to Calculus","teacher_id":1}'
```

Create an enrollment:
```bash
curl -X POST http://localhost:8080/enrollments \
  -H "Content-Type: application/json" \
  -d '{"student_id":1,"course_id":1,"grade":"A"}'
```


