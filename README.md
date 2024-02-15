# technical-test-YC-W22

This project is structured based on the principles of Clean Architecture to promote separation of concerns and maintainability. The codebase is organized into layers, each with distinct responsibilities, facilitating testability and scalability.

.
├── cmd             # Command-line application entry point
│   └── yourapp
├── entity          # Business entities
├── handlers        # HTTP request handlers
├── model           # Presentation layer models
├── repository      # Data storage and retrieval
│   └── postgres    # PostgreSQL repository implementation
├── service         # Business logic
├── migrations      # Database migration files
└── utils           # Utility functions


We use existing libs :


- Mux Router

- Godotenv, for env loader

- Gorm, for ORM

## Setups
-  Clone Repo
```
- after clone repo run
$ 	go mod tidy
   ```
- Install [golang-migrate](https://github.com/golang-migrate/migrate?tab=readme-ov-file) for migration tools
- Setup .env
- [optional] Run sample migrations
```
$ 	migrate -path ./migrations -database "postgres://postgres@localhost:5432/technical-test?sslmode=disable" up
```
 - Running Application
 ```
$ go run main.go
 ```
- to runnning unit test (for now only service applied UT)
```
$ go test ./service -v
