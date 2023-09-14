package bootstrap

import (
	"embed"
	"os"
)

// Env : Configuration struct that holds all the environment variables used for the application
type Env struct {
	ServerAddress string
	DatabasePath  string
}

// NewEnv: create a new instance for Env
func NewEnv(fs embed.FS) *Env {
	env := Env{
		ServerAddress: getEnvOrDefault("SERVER_ADDRESS", ":8080"),
		DatabasePath:  getEnvOrDefault("DATABASE_PATH", "data/whoknowkmh-portfolio.db"),
	}

	return &env
}

// getEnvOrDefault: Look up a variable's value using its specific variable name in configuration,
// If the lookup for the variable fails, set its value to the defaultValue we specified earlier
func getEnvOrDefault(name, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		value = defaultValue
	}
	return value
}
