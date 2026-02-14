package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	HttpServerConfig HttpServerConfig
	Database         Database
	RedisAddr        string
}

type HttpServerConfig struct {
	ListenAddress string
}

type Database struct {
	Name       string
	Schema     string
	Hosts      string
	User       string
	UserSlaves string
	password   string
	Port       int
	SSLMode    string
}

func (d *Database) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		d.Hosts,
		d.Port,
		d.User,
		d.Name,
		d.password,
		d.SSLMode,
		d.Schema,
	)
}

func NewDefaultConfig() *Config {
	return &Config{
		HttpServerConfig: HttpServerConfig{
			ListenAddress: getEnv("APP_HTTP_ADDR", "0.0.0.0:10080"),
		},
		Database: Database{
			Hosts:    getEnv("APP_DB_HOST", "localhost"),
			Port:     getEnvAsInt("APP_DB_PORT", 8432),
			User:     getEnv("APP_DB_USER", "postgres"),
			Name:     getEnv("APP_DB_NAME", "postgres"),
			password: getEnv("APP_DB_PASSWORD", "postgres"),
			SSLMode:  getEnv("APP_DB_SSLMODE", "disable"),
			Schema:   getEnv("APP_DB_SCHEMA", "public"),
		},
		RedisAddr: getEnv("APP_REDIS_ADDR", ":8379"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
