package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"

	"github.com/TomChv/csc-0847/project_1/backend/utils"
)

const configPath = "envs/backend.env"

var (
	corsOrigins    []string
	allowAllOrigin bool
)

func loadConfig() {
	// Ignore dotenv loading on CI
	if os.Getenv("CI") == "false" {
		if err := godotenv.Load(configPath); err != nil {
			log.Fatalln(err)
		}
	}

	allowAllOrigin = utils.ForceGetEnv("BACKEND_CORS_ALL") == "true"

	if !allowAllOrigin {
		corsOrigins = strings.Split(utils.ForceGetEnv("BACKEND_CORS"), ",")
	}
}
