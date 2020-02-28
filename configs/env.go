package configs

import (
	"os"
	"time"
)

func getEnv(key, fallback string) string {
	s := os.Getenv(key)
	if len(s) == 0 {
		return fallback
	}
	return s
}

var (
	HttpTimeout = time.Second * 15
	// mongodb configs
	MongoURI    = getEnv("MONGO_URI", "mongodb://127.0.0.1:27018")
	MongoDbName = getEnv("MONGO_DB_NAME", "users-service")
	// application port
	AppPort = getEnv("APP_PORT", "5000")
)
