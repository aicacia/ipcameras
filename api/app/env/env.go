package env

import (
	"os"
)

func IsProd() bool {
	return os.Getenv("APP_ENV") == "prod"
}

func IsDev() bool {
	return os.Getenv("APP_ENV") == "dev"
}

func IsTest() bool {
	return os.Getenv("APP_ENV") == "test"
}

func GetDatabaseUrl() string {
	return os.Getenv("DATABASE_URL")
}
