# AhuM

AhuM is a command-line interface (CLI) tool designed to streamline the development and management of microservices. It helps in initializing project structures, configuring infrastructures, and managing various aspects of microservice development, including service creation, edge management, and database interactions.

## Installation

To install AhuM, use the following Go command:

```bash
go install github.com/Ahu-Tools/AhuM@latest
```

## Milestones

Here's a breakdown of the project's development milestones:

### 1. Initialise an Ahu project structure ✅
    - 1.1. Basic directories ✅
    - 1.2. Basic config.json generation ✅
    - 1.3. Basic config.go generation ✅
    - 1.4. Go mod init at the end ✅
    - 1.5. Go mod tidy at the end ✅

### 2. Initialising infrastructures ✅
    - 2.1. Form list to choose ✅
    - 2.2. Run registered infrastructures forms ✅
    - 2.3. Inject infrastructures configurations to config.json ✅
    - 2.4. Inject infrastructures configurations loadings to config.go ✅
    - 2.5. Basic infrastructures files generations ✅

### 3. Create an Ahu service (chain, service, data) ✅

### 4. Create an Ahu Edge ✅
    - 4.1. Gin (default) ✅
    - 4.2. Connect ✅
        - 4.1.2. Generate Connect files ✅

### 5. Kafka

### 6. Postgres Repo Management
    - 6.1. Gorm GEN
        - 6.1.1. Database to struct
        - 6.1.2. Dynamic SQL
        - 6.1.3. DAO Generation
    - 6.2. Goose Migration
        - 6.2.1. Migration creation
        - 6.2.2. Port goose commands

### 7. Test Management
