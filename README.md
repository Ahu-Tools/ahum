# AhuM

<p align="center">
  <strong>A powerful CLI tool for building Go microservices with Hexagonal Architecture</strong>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#commands">Commands</a> •
  <a href="#architecture">Architecture</a> •
  <a href="#roadmap">Roadmap</a>
</p>

---

AhuM is a command-line interface (CLI) tool designed to streamline the development and management of Go microservices. It automates project scaffolding, code generation, and maintains consistency across the codebase following **Hexagonal Architecture** (Ports & Adapters) principles.

## Features

- 🚀 **Interactive Project Initialization** - TUI-based project setup wizard
- 📦 **Service Generation** - Automated creation of chain, service, and data layers
- 🔌 **Multiple Edge Support** - Connect (gRPC), Gin (REST), Asynq (background tasks)
- 🏗️ **Infrastructure Management** - PostgreSQL, Redis, and more
- 🔄 **Code Injection** - Safe code modifications using marker comments
- 📝 **Protobuf Integration** - Automatic `buf generate` execution

## Installation

```bash
go install github.com/Ahu-Tools/ahum@latest
```

Verify installation:

```bash
ahum --help
```

## Quick Start

### Create a New Project

```bash
# Initialize a new project with interactive wizard
ahum init
```

The TUI will guide you through:
1. Project name and Go module path
2. Infrastructure selection (PostgreSQL, Redis, etc.)
3. Edge selection (Connect, Gin, Asynq)
4. Initial configuration

### Add a Feature

```bash
# Create a new service (chain, service, data layers)
ahum service create

# Add a Connect service
ahum connect service add product

# Add a version
ahum connect service version add v1 product

# Add methods
ahum connect service method add CreateProduct product v1
ahum connect service method add GetProduct product v1
```

## Commands

### Project Management

| Command | Description |
|---------|-------------|
| `ahum init` | Initialize a new Ahu project with interactive TUI |

### Service Management

| Command | Description |
|---------|-------------|
| `ahum service create` | Create a new service with chain, service, and data layers |

### Infrastructure Management

| Command | Description |
|---------|-------------|
| `ahum infra generate` | Add new infrastructure components via interactive form |

### Edge Management

| Command | Description |
|---------|-------------|
| `ahum edge generate` | Generate a new edge (Connect, Gin, Asynq) |

### Connect (gRPC) Commands

| Command | Description |
|---------|-------------|
| `ahum connect service add [name]` | Add a new Connect service |
| `ahum connect service version add [version] [service]` | Add a version to a service |
| `ahum connect service method add [method] [service] [version]` | Add a method to a service version |
| `ahum connect gen` | Regenerate protobuf code with `buf generate` |

### Gin (REST) Commands

| Command | Description |
|---------|-------------|
| `ahum gin route version add [version]` | Add a new API version |
| `ahum gin route entity add [version] [entity]` | Add an entity to a version |
| `ahum gin route handle add [version] [entity] [method]` | Add a handler method |

### Asynq (Background Tasks) Commands

| Command | Description |
|---------|-------------|
| `ahum asynq edge add module [version] [module]` | Add a task module |
| `ahum asynq edge add task [version] [module] [task]` | Add a task handler |

## Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--path` | `-p` | Project root path (default: current directory) |
| `--config` | | Custom config file path (default: `$HOME/.ahum.yaml`) |

**Example:**
```bash
ahum -p /path/to/project connect service add user
```

## Architecture

AhuM generates projects following **Hexagonal Architecture**:

```
┌─────────────────────────────────────────────────────────────────┐
│                         EDGES (Inbound)                         │
│   Connect (gRPC)  │  Gin (REST)  │  Asynq (Tasks)  │  Events   │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    CHAIN (Application Layer)                    │
│      Dependency Injection  │  Validation  │  Error Mapping     │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    SERVICE (Business Logic)                     │
│         Use Cases  │  Business Rules  │  Domain Operations     │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                      DATA (Domain Layer)                        │
│      Entities  │  Repository Interfaces  │  Domain Contracts   │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                 INFRASTRUCTURE (Outbound Adapters)              │
│      PostgreSQL  │  Redis  │  SMS  │  Email  │  External APIs  │
└─────────────────────────────────────────────────────────────────┘
```

### Generated Project Structure

```
.
├── main.go                     # Application entry point
├── config/
│   ├── config.go               # Configuration loading
│   └── config.json             # Application configuration
├── edge/
│   ├── edge.go                 # Edge interface
│   └── connect/                # Connect edge (or gin/, asynq/)
│       ├── connect.go          # Server setup
│       └── {service}/          # Service handlers
│           ├── registrar.go
│           └── v1/
│               ├── edge.go
│               └── *.proto
├── chain/
│   └── {service}/
│       ├── chain.go            # DI composition
│       ├── entity.go           # DTOs
│       └── handler.go          # Orchestration
├── service/
│   └── {service}/
│       ├── service.go          # Constructor
│       ├── entity.go           # DTOs
│       └── handler.go          # Business logic
├── data/
│   └── {service}/
│       ├── entity.go           # Domain entities
│       └── repo.go             # Repository interface
└── infrastructure/
    └── postgres/
        └── {service}/
            ├── model.go        # DB models
            └── repo.go         # Implementation
```

## Code Markers

AhuM uses special comment markers to identify injection points for safe code modifications:

```go
// edge/connect/connect.go
func RegisterServices(mux *http.ServeMux) {
    hello.RegisterService(mux)
    // @ahum: services    <-- New services injected here
}
```

```go
// config/config.go
func ConfigInfras() error {
    // @ahum:infras.group
    err := postgres.Configure()
    // @ahum:end.infras.group
    
    //@ahum: loads
    return nil
}
```

### Available Markers

| Marker | Location | Purpose |
|--------|----------|---------|
| `// @ahum: imports` | File imports | Import statements injection |
| `// @ahum: services` | connect.go | Service registration |
| `// @ahum: versions` | registrar.go | Version registration |
| `// @ahum: methods` | edge.go | Method stubs |
| `// @ahum: edges` | edge/edge.go | Edge registration |
| `// @ahum: loads` | config.go | Infrastructure loading |
| `// @ahum:infras.group` | config.go | Infrastructure group block |

## Configuration

AhuM stores global configuration in `$HOME/.ahum.yaml`:

```yaml
# Default project path
projectPath: .

# Custom templates path (optional)
templatesPath: ~/.ahum/templates
```

### Project Detection

AhuM automatically detects project metadata by analyzing:

1. **Directory Structure** - Recognizes the hexagonal architecture layout (`edge/`, `chain/`, `service/`, `data/`, `infrastructure/`)
2. **Code Analysis** - Reads existing code to understand registered services, versions, and methods
3. **Configuration File** - Parses `config/config.json` for infrastructure and edge configurations

```json
// config/config.json - AhuM reads this for project context
{
  "app": {
    "secret_key": "...",
    "env": "dev"
  },
  "edges": {
    "connect": {
      "server": { "host": "0.0.0.0", "port": "8080" }
    }
  },
  "infras": {
    "postgres": {
      "user": "postgres",
      "host": "localhost",
      "port": "5432"
    }
  }
}
```

This approach means:
- **No additional metadata files** required
- **Works with existing projects** that follow the architecture
- **Stays in sync** with actual code state

## Roadmap

### ✅ Completed

- [x] **Project Initialization** - Interactive TUI wizard
- [x] **Service Generation** - Chain, service, data layers
- [x] **Infrastructure Management** - PostgreSQL, Redis support
- [x] **Connect Edge** - gRPC-Web with service/version/method management
- [x] **Gin Edge** - REST API with version/entity/handler management
- [x] **Asynq Edge** - Background task processing

### 📋 Planned

- [ ] **Kafka Integration** - Event-driven messaging

- [ ] **PostgreSQL Repository Management**
  - [ ] GORM GEN integration (DB to struct, dynamic SQL, DAO)
  - [ ] Goose migration management

- [ ] **Test Management**
  - [ ] Automated test scaffolding
  - [ ] Mock generation

- [ ] **Security Package**
  - [ ] SecureString type generation
  - [ ] GORM encryption hooks
  - [ ] JSON redaction utilities

- [ ] **Crypto Package**
  - [ ] Encrypter interface scaffolding
  - [ ] AES-GCM implementation
  - [ ] Blind index computation

- [ ] **JWT Helper Package**
  - [ ] RSA JWT generation utilities
  - [ ] Token validation helpers
  - [ ] Configuration scaffolding

- [ ] **Logging Package**
  - [ ] Structured slog setup
  - [ ] JSON handler configuration

- [ ] **DevOps & CI/CD**
  - [ ] Dockerfile generation (multi-stage builds)
  - [ ] .dockerignore generation
  - [ ] Docker Compose templates
  - [ ] GitHub Actions workflows
  - [ ] GitLab CI/CD pipelines
  - [ ] Kubernetes manifests
  - [ ] Helm chart scaffolding
  - [ ] Makefile generation

## Examples

### Complete Feature Workflow

```bash
# 1. Initialize project
ahum init

# 2. Create a "product" service
ahum service create
# Enter: product

# 3. Add Connect service and endpoints
ahum connect service add product
ahum connect service version add v1 product
ahum connect service method add CreateProduct product v1
ahum connect service method add GetProduct product v1
ahum connect service method add UpdateProduct product v1
ahum connect service method add DeleteProduct product v1
ahum connect service method add ListProducts product v1

# 4. Add infrastructure
ahum infra generate
# Select: PostgreSQL

# 5. Regenerate protos after manual edits
ahum connect gen
```

### Adding a New API Version

```bash
# Add v2 to existing product service
ahum connect service version add v2 product

# Add new methods to v2
ahum connect service method add GetProductDetails product v2
```

### Adding Gin REST Endpoints

```bash
# Add version
ahum gin route version add v1

# Add entity
ahum gin route entity add v1 product

# Add handlers
ahum gin route handle add v1 product Create
ahum gin route handle add v1 product Get
ahum gin route handle add v1 product List
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Sina Sadeghi** - [sina.sadeghi83@gmail.com](mailto:sina.sadeghi83@gmail.com)

---

<p align="center">
  Made with ❤️ for the Go community
</p>