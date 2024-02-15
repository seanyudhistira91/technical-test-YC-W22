package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type SqlDbConf struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	SSLMode                string
	Schema                 string
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnSeconds int64
	MaxLifeTimeConnSeconds int64
}

type LogConf struct {
	Name string
}

// Config ...
type Config struct {
	App          AppConf
	SqlDb        SqlDbConf
	Http         HttpConf
	Log          LogConf
	Debug        bool
	IsLogEnabled bool
}

var osGetenv = os.Getenv

// Make builds a config value based on .env file.
func Make() Config {
	app := AppConf{
		Environment: osGetenv("APP_ENV"),
		Name:        osGetenv("APP_NAME"),
	}

	sqldb := SqlDbConf{
		Host:     osGetenv("DB_HOST"),
		Username: osGetenv("DB_USERNAME"),
		Password: osGetenv("DB_PASSWORD"),
		Name:     osGetenv("DB_NAME"),
		Port:     osGetenv("DB_PORT"),
		SSLMode:  osGetenv("DB_SSL_MODE"),
		Schema:   osGetenv("DB_SCHEMA"),
	}

	http := HttpConf{
		Port: osGetenv("HTTP_PORT"),
	}

	log := LogConf{
		Name: osGetenv("LOG_NAME"),
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(osGetenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	dBMaxOpenConn, err := strconv.Atoi(osGetenv("DB_MAX_OPEN_CONN"))
	if err == nil {
		sqldb.MaxOpenConn = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(osGetenv("DB_MAX_IDLE_CONN"))
	if err == nil {
		sqldb.MaxIdleConn = dBMaxIdleConn
	}

	dBMaxIdleTimeConnSeconds, err := strconv.Atoi(osGetenv("DB_MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		sqldb.MaxIdleTimeConnSeconds = int64(dBMaxIdleTimeConnSeconds)
	}

	dBMaxLifeTimeConnSeconds, err := strconv.Atoi(osGetenv("DB_MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		sqldb.MaxLifeTimeConnSeconds = int64(dBMaxLifeTimeConnSeconds)
	}
	debug, _ := strconv.ParseBool(osGetenv("DEBUG"))
	logEnabled, _ := strconv.ParseBool(osGetenv("LOG_ENABLED"))

	config := Config{
		App:          app,
		SqlDb:        sqldb,
		Http:         http,
		Log:          log,
		Debug:        debug,
		IsLogEnabled: logEnabled,
	}

	return config
}
