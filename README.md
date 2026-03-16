# trae-golang

Sample API app using Go + Gin + SQLite (GORM), with JWT authentication, standard JSON responses, error samples, Swagger docs, and OpenTelemetry export.

## Run

```bash
go mod tidy
go run main.go
```

- API base: `http://localhost:8080/api/v1`
- Swagger UI: `http://localhost:8080/docs-api/index.html`

## OpenTelemetry

This app exports traces and metrics using OTLP/HTTP.

Default exporter endpoint:

- `OTEL_EXPORTER_OTLP_ENDPOINT=http://172.15.1.100:4318`

Optional env vars:

- `OTEL_SERVICE_NAME=sample-api`
- `OTEL_EXPORTER_OTLP_ENDPOINT=http://<collector-host>:4318`

Example:

```bash
export OTEL_SERVICE_NAME="sample-api"
export OTEL_EXPORTER_OTLP_ENDPOINT="http://172.15.1.100:4318"
go run main.go
```

## Auth (JWT)

Register:

```bash
curl -s -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```

Login:

```bash
curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

Use token (example):

```bash
curl -s http://localhost:8080/api/v1/user/profile \
  -H "Authorization: Bearer <token>"
```
