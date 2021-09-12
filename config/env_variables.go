package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Env                      string
	Port                     int
	AccessControlAllowOrigin string
	DatabaseUrl              string
	DeployHooksUrl           string
}

func NewEnvVariables() *EnvVariables {
	return &EnvVariables{}
}

func (e *EnvVariables) Init() error {
	err := godotenv.Load(".env")

	e.Env = os.Getenv("ENV")
	isLocalEnv := e.Env == "local"

	if err != nil {
		if isLocalEnv {
			return fmt.Errorf("%s", err)
		}
	} else {
		if !isLocalEnv {
			return fmt.Errorf("%s", err)
		}
	}

	parsedPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	e.Port = parsedPort
	e.AccessControlAllowOrigin = os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN")
	e.DatabaseUrl = os.Getenv("DATABASE_URL")
	e.DeployHooksUrl = os.Getenv("DEPLOY_HOOKS_URL")

	return nil
}
