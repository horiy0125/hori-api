package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Env                      string
	Port                     string
	AccessControlAllowOrigin string
	DatabaseUrl              string
}

func NewEnvVariables() *EnvVariables {
	return &EnvVariables{}
}

func (e *EnvVariables) Init() {
	err := godotenv.Load(".env")

	e.Env = os.Getenv("ENV")
	isLocalEnv := e.Env == "local"

	if err != nil {
		if isLocalEnv {
			panic("")
		}
	} else {
		if !isLocalEnv {
			panic("")
		}
	}

	// parsedPort, err := strconv.Atoi(os.Getenv("PORT"))
	// if err != nil {
	// 	panic("")
	// }

	e.Port = os.Getenv("PORT")
	e.AccessControlAllowOrigin = os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN")
	e.DatabaseUrl = os.Getenv("DATABASE_URL")
}
