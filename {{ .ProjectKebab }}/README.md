# {{ .Project }}

{{ .Scaffold.description }}

## Overview

This project consists of:
- Go backend API
- React frontend (built with Vite and Bun)

## Development

### Prerequisites
- Go 1.21+
- Bun 1.0+ (https://bun.sh/)
- Docker and Docker Compose (optional)

### Using Docker

The simplest way to start the application is using Docker Compose:

```bash
docker-compose up
```

This will start both the backend and frontend services:
- Backend: http://localhost:8080
- Frontend: http://localhost:5173

### Manual Setup

#### Backend
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

#### Frontend
```bash
cd frontend
bun install
bun run dev
```

## Project Structure

```
.
├── backend/          # Go backend API
│   ├── cmd/          # Application entrypoints
│   └── internal/     # Private application code
├── frontend/         # React frontend
│   └── src/          # Frontend source code
```

## CI/CD

This project uses GitHub Actions for continuous integration and deployment. The workflow:
1. Builds and tests the backend
2. Builds and tests the frontend
3. Creates Docker images (on tag/release)

## Author

{{ .Scaffold.author }} <{{ .Scaffold.email }}>
