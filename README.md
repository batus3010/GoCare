# GoCare: Clinic Management Application

A simple Golang web application that provides two user portals:

- **Receptionist Portal**: Register and manage patients (CRUD).
- **Doctor Portal**: View and update patient details.

A single authentication API serves both roles. Built with Clean Architecture principles (transport → business → storage).

---

## Technology Stack

- **Language**: Go (1.20+)
- **Web Framework**: Gin
- **ORM**: GORM (PostgreSQL)
- **Authentication**: JWT tokens
- **Testing**: `testing` package (unit tests), `httptest` (integration)

---

## Features

- **User Management**
    - Registration (receptionist role only)
    - Login (JWT authentication)
- **Role-Based Access Control**
    - Receptionists: Create, Read, Update, Delete patients
    - Doctors: Read and Update patients
- **Patient Management**
    - Soft-delete patients with status flag
    - Pagination support on patient listing
- **Clean Architecture**
    - `components/` — shared utilities (app context, middleware, hasher, tokenprovider)
    - `module/user/`, `module/patient/` — each with `biz/`, `model/`, `storage/`, `transport/gin/`
- **Testing**
    - Business layer unit tests
    - Repository integration tests
    - HTTP transport tests (using Gin + `httptest`)

---

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL

### Environment Variables

Create a `.env` file or export:

```bash
DB_CONN_STR="host=... user=... password=... dbname=... sslmode=disable"
SYSTEM_SECRET="your_jwt_secret"
HTTP_ADDR=":8080"
```

### Run the server
go run main.go

The server will listen on `http://localhost:8080` by default.
### Testing
go run ./...

### Project structure

```text
GoCare/                      # root project folder
├── common/                  # common constants, errors, types
├── components/              # shared infrastructure
│   ├── appctx/              # application context (DB, config)
│   ├── hasher/              # password hashing
│   ├── tokenprovider/       # JWT provider interface and implementation
│   └── middleware/          # Gin middleware (auth, recovery)
├── module/                  # feature modules
│   ├── user/                # user management (auth, register)
│   │   ├── biz/             # business logic & unit tests
│   │   ├── model/           # data models & DTOs
│   │   ├── storage/         # database storage & integration tests
│   │   └── transport/gin/   # HTTP handlers & transport tests
│   └── patient/             # patient CRUD module
│       ├── biz/
│       ├── model/
│       ├── storage/
│       └── transport/gin/
├── cmd/                     # optional entrypoints if using cmd structure
│   └── server/              # server main.go
├── main.go                  # application bootstrap (alternative entry)
├── go.mod                   # module definition
└── README.md                # project documentation