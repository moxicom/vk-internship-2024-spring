package models

type PostgresConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresName     string
	PostgresSSLMode  string
}
