# SysTrace Server

A simple Go backend for system/device telemetry with WebSocket updates and PostgreSQL support.

> Note: This is the **Server component** of SysTrace. You also need the **SysTrace Agent** to collect telemetry data.

## Demo

- [Demo Video](https://youtu.be/qxP3Q3q_ucY)

## Requirements

- Go 1.26+
- PostgreSQL 15+
- Docker Desktop installed
- Docker Engine running (`docker info` must succeed)

## Installer Prerequisites

The installer expects Docker to be available before setup starts:

- `docker.exe` must be available in `PATH`
- Docker Desktop must be started
- Docker Engine must be reachable (`docker info` exits with code 0)

If this is not fulfilled, the installer aborts early with a message.

## Quick Start

1. Install dependencies:

```bash
go mod tidy
```

2. Run the server:

```bash
go run .
```

## Project Structure (Short)

- `main.go` - application entry point
- `services/web/` - HTTP wiring and server startup
- `services/handler/` - HTTP and WebSocket handlers
- `services/database/` - database connection and logic
- `ws/` - WebSocket hub, client, events, and responses
- `data/static/` - static data models
- `templates/` - frontend assets (`app.js`, `style.css`)
- `db/` - SQL schema and Docker Compose files

## Related Repositories

- SysTrace Agent: https://github.com/EliasSchaffer/SysTrace_Agent.git
