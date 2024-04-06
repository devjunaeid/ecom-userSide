package config

import (
	"github.com/joho/godotenv"
	"os"
)

type myEnvs struct {
	DbConnectionUrl string
}

var Envs = initEnvs()

func initEnvs() *myEnvs {
	godotenv.Load()
	return &myEnvs{
		DbConnectionUrl: os.Getenv("PSQL_CONN_URL"),
	}
}
