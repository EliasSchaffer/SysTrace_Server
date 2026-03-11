# SysTrace Server

A simple Go backend for system/device telemetry with WebSocket updates and PostgreSQL support.

## Requirements

- Go 1.26+
- PostgreSQL (optional for features that require persistence)

## Quick Start

1. Install dependencies:

```bash
go mod tidy
```

2. Run the server:

```bash
go run .
```

## Project Structure

- `main.go` - application entrypoint
- `web/` - web server startup and HTTP wiring
- `services/handler/` - HTTP and WebSocket handlers
- `services/database/` - database connection and logic
- `ws/` - WebSocket hub, client, events, and responses
- `data/static/` - static data models (CPU, RAM, GPS, hardware, device)
- `templates/` - frontend templates and static assets (`app.js`, `style.css`)
- `db/` - SQL schema and Docker Compose for database setup

## Database Notes

SQL scripts and a compose file are provided in `db/`:

- `create_tables.sql`
- `optimized_schema.sql`
- `docker-compose.yml`

Use these if you want to run with a local PostgreSQL instance.

## Related Repositories

- SysTrace Agent: https://github.com/EliasSchaffer/SysTrace_Agent.git
