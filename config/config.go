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
	Port             = 0
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5432
	}

	ConnectionString = fmt.Sprintf("postgres://postgres:postgres@localhost:%d/postgres?sslmode=disable", Port)
}
