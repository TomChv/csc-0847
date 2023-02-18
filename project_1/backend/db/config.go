package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/TomChv/csc-0847/project_1/backend/utils"
)

const configPath = "envs/database.env"

var (
	host     string
	port     string
	name     string
	user     string
	password string
	provider Provider
)

func loadConfig() {
	// Ignore dotenv loading on CI
	if os.Getenv("CI") == "false" {
		if err := godotenv.Load(configPath); err != nil {
			log.Fatalln(err)
		}
	}

	// Ignore those variables during unit tests
	// because we are using a SQLite3 database.
	if os.Getenv("UNIT_TEST") == "" {
		host = utils.ForceGetEnv("DB_HOST")
		port = utils.ForceGetEnv("DB_PORT")
		name = utils.ForceGetEnv("DB_NAME")
		user = utils.ForceGetEnv("DB_USER")
		password = utils.ForceGetEnv("DB_PASSWORD")
	}

	_provider, err := stringToProvider(utils.ForceGetEnv("DB_PROVIDER"))
	if err != nil {
		log.Fatalln(err)
	}

	provider = _provider
}
