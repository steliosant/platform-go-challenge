# Platform Go Challenge API

REST API for users, assets, and favourites with JWT authentication. Ships with Docker, Swagger UI, and seed data so you can try it immediately.

## Stack
- Go 1.25.6
- PostgreSQL 15
- Docker + docker-compose
- JWT via github.com/golang-jwt/jwt/v5
- Password hashing via golang.org/x/crypto/bcrypt
- Swagger UI via github.com/swaggo/http-swagger

## Quickstart (Docker)
```bash
cd docker
docker compose up --build
```
- API: http://localhost:8000
- Swagger UI: http://localhost:8000/swagger/index.html
- Adminer (DB UI): http://localhost:8080 (system=PostgreSQL, server=db, user=password as below)

## Quickstart (local without Docker)
1) Start PostgreSQL and create the database:
```bash
createdb dashboard
psql dashboard -c "CREATE USER user WITH PASSWORD 'password';"
psql dashboard -c "GRANT ALL PRIVILEGES ON DATABASE dashboard TO user;"
psql dashboard -f db/init/001_init.sql
```
2) Run the API:
```bash
DATABASE_URL=postgres://user:password@localhost:5432/dashboard?sslmode=disable PORT=8080 go run main.go
```
API will be at http://localhost:8080.

## Authentication
- JWT bearer tokens, signed with HS256.
- Configure secret with env `JWT_SECRET` (defaults to a dev fallback).
- To disable auth for testing, set `AUTH_DISABLED=true` (all protected endpoints become open).

### Seeded users
- u1 / alice123
- u2 / bob123

### Auth endpoints
- POST /register — create user (id, name, password)
- POST /login — returns JWT token

Include the token in `Authorization: Bearer <token>` for protected routes.

## Endpoints (summary)
- Auth: POST /login, POST /register
- Users: GET /users, POST /users
- Favourites: GET /users/{userId}/favourites, POST /users/{userId}/favourites, PATCH /users/{userId}/favourites/{assetId}, DELETE /users/{userId}/favourites/{assetId}
- Assets: GET /assets/{id}, POST /assets

See full request/response schemas in Swagger UI.

## Database seeding
`db/init/001_init.sql` creates tables and seeds:
- Users u1/u2 (with bcrypt password hashes)
- Sample assets (insight, chart)
- Sample favourites linking users to assets

## Running tests
```bash
go test ./...
```

## Configuration
- `DATABASE_URL` (required): postgres connection string.
- `PORT` (default 8080): HTTP port inside the container/process.
- `JWT_SECRET`: secret for signing tokens.
- `AUTH_DISABLED`: set to `true` to bypass auth checks (use only for local testing).

## Useful URLs (Docker defaults)
- API base: http://localhost:8000
- Swagger: http://localhost:8000/swagger/index.html
- Adminer: http://localhost:8080 (server=db, user=user, password=password, database=dashboard)
