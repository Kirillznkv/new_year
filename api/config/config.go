package config

import "os"

var (
	PostgresDB       = os.Getenv("POSTGRES_DB")
	PostgresUser     = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	DatabaseURL      = os.Getenv("DATABASE_URL")
)
