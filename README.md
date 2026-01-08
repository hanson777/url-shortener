# URL Shortener

A simple URL shortener service built with Go and Svelte.

## Prerequisites

- Go 1.21+
- Node.js and npm
- PostgreSQL (use pgAdmin or any other method to run a local Postgres database)

## Quick Start

### Database

Create a `.env` file in the root directory with your Postgres connection string:

```
POSTGRES_URL=postgres://username:password@localhost:5432/dbname
```

### Backend

Start the backend server:

```bash
go run cmd/api/main.go
```

### Frontend

Navigate to the frontend directory and start the development server:

```bash
cd frontend
npm install  # First time only
npm run dev
```

The frontend will be available at `http://localhost:5173` (or the port shown in the terminal).
