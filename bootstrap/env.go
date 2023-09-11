package bootstrap

import (
	"embed"
	"os"
)

type Env struct {
	ServerAddress string
	DatabasePath  string
}

func NewEnv(fs embed.FS) *Env {
	env := Env{
		ServerAddress: getEnvOrDefault("SERVER_ADDRESS", ":8080"),
		DatabasePath:  getEnvOrDefault("DATABASE_PATH", "data/whoknowkmh-portfolio.db"),
	}

	return &env
}

func getEnvOrDefault(name, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		value = defaultValue
	}
	return value
}
