// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/seanyudhistira91/technical-test-YC-W22/config"
	"github.com/seanyudhistira91/technical-test-YC-W22/handlers"
	"github.com/seanyudhistira91/technical-test-YC-W22/repository"
	"github.com/seanyudhistira91/technical-test-YC-W22/repository/postgres"
	"github.com/seanyudhistira91/technical-test-YC-W22/service"
	"github.com/sirupsen/logrus"
)

var postgresNew = postgres.New
var configMake = config.Make

func main() {
	// read the server environment variables
	conf := configMake()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// open connection to db
	postgresdb := postgresNew(conf.SqlDb, logrus.New(), isProd)
	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logrus.New(), postgresdb.SqlDB, postgresdb.DB.Name())

	// Initialize repository, service, and handler
	repo := repository.NewRepository(postgresdb.DB)
	svc := service.NewService(repo)
	handler := handlers.NewHandler(svc)

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/api/registration", handler.Registration).Methods("POST")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")

	// Start HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Http.Port), router))
}
