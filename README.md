## Features

- Modular Routers: Organized under /api, cleanly separated.
- GORM: Includes models for Book and AuditLog.
- CRUD APIs: Supports Create, Read, Update, and Delete for Book
- Audit Log (Event): Automatically logs events after inserts using event listeners.
- Middlewares: Easy to integrate logging, authentication, or performance profiling.

## Quick Start

Go 1.21.1 or later (If not already installed) in your machine.

### 1. Clone the Repo

```
git clone <repo-url>
cd go-ginapi-starter
```
### 2. Install pkg and Run app

```
go mod tidy
go run main.go
```
