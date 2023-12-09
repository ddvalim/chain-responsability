package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ConnectionString = ""
	APIPort          = 0
	DBPort           = 0
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 8686
	}

	DBPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		DBPort = 5432
	}

	ConnectionString = fmt.Sprintf("postgres://postgres:postgres@localhost:%d/postgres?sslmode=disable", DBPort)
}
