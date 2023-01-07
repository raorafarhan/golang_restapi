package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	SERVER_PORT int
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()

	}

	return appConfig
}

func initConfig() *AppConfig {
	var defConfig AppConfig
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	defConfig.DB_USERNAME = os.Getenv("MYSQL_USER")
	defConfig.DB_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	defConfig.DB_HOST = os.Getenv("MYSQL_HOST")
	defConfig.DB_NAME = os.Getenv("MYSQL_DBNAME")

	dbPortConv, errDBPort := strconv.Atoi(os.Getenv("MYSQL_PORT"))

	if errDBPort != nil {
		log.Fatal(errDBPort)
		return nil
	}

	defConfig.DB_PORT = dbPortConv

	serverPortConv, errServerPort := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if errServerPort != nil {
		log.Fatal(errServerPort)
		return nil
	}

	defConfig.SERVER_PORT = serverPortConv

	return &defConfig
}
