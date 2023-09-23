package bootstrap

import (
	"embed"
	"os"
	"strconv"
)

// Env : Configuration struct that holds all the environment variables used for the application
type Env struct {
	ServerAddress              string
	DatabasePath               string
	SessionSecret              string
	AccessTokenSecret          string
	AccessTokenExpiryHour      int
	CacheStaticFilesExpiryHour int
	CSRFAuthKey                string // 32 byte key
	AdminUserID                string
	AdminUserPwdHash           string // https://go.dev/play/p/uKMMCzJWGsW
}

// NewEnv: create a new instance for Env
func NewEnv(fs embed.FS) *Env {
	env := Env{
		ServerAddress:              getEnvOrDefault("SERVER_ADDRESS", ":8080"),
		DatabasePath:               getEnvOrDefault("DATABASE_PATH", "data/whoknowkmh-portfolio.db"),
		SessionSecret:              getEnvOrDefault("SECRET", "whoknowkmh-secret"),
		AccessTokenSecret:          getEnvOrDefault("ACCESS_TOKEN_SECRET", "whoknowkmh-secret"),
		AccessTokenExpiryHour:      stringToInt(getEnvOrDefault("ACCESS_TOKEN_EXPIRY_HOUR ", "2")),
		CacheStaticFilesExpiryHour: stringToInt(getEnvOrDefault("CACHE_STATIC_FILES_EXPIRY_HOUR ", "2")),
		CSRFAuthKey:                getEnvOrDefault("CSRF_AUTH_KEY ", "0d3836af82a643f7b4212fb4c125ef5c8146310a4fec68fa36d2f86d89eeccf7032eb0ad1183d22bb382741030685a7dc909a09f09006e290dc45ca93ac4bc26"),
		AdminUserID:                getEnvOrDefault("ADMIN_USER_ID ", "admin"),
		AdminUserPwdHash:           getEnvOrDefault("ADMIN_USER_PWD ", "$2a$12$RPAJZXqgm2bP/VIWsJUUcuNSG57/EzbMvD0ebch3e1q518oGaSSgu"), // pwd default 123456
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

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
