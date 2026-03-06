package config

import "os"

type Config struct {
	GRPC     GRPCConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type GRPCConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func LoadConfig() *Config {
	return &Config{
		GRPC: GRPCConfig{
			Port: getEn("GRPC_PORT", "50051"),
		},
		Postgres: PostgresConfig{
			Host:     getEn("POSTGRES_HOST", "localhost"),
			Port:     getEn("POSTGRES_PORT", "5432"),
			User:     getEn("POSTGRES_USER", "postgres"),
			Password: getEn("POSTGRES_PASSWORD", ""),
			DBName:   getEn("POSTGRES_DB", "fm_user_db"),
		},
		Redis: RedisConfig{
			Host: getEn("REDIS_HOST", "localhost"),
			Port: getEn("REDIS_PORT", "6379"),
		},
	}
}

func getEn(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}
