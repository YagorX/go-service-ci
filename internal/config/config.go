package config

import (
	"fmt"
)

type Config struct {
	HttpServerConfig HttpServerConfig
	Database         Database
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
	Password   string
	Port       int
	SSLMode    string
}

func (d *Database) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		d.Hosts,
		d.Port,
		d.User,
		d.Name,
		d.Password,
		d.SSLMode,
		d.Schema,
	)
}

func NewDefaultConfig() *Config {
	return &Config{
		HttpServerConfig: HttpServerConfig{
			ListenAddress: "0.0.0.0:10080",
		},
		Database: Database{
			Hosts:    "localhost",
			Port:     8432,
			User:     "postgres",
			Name:     "postgres",
			Password: "postgres",
			SSLMode:  "disable",
			Schema:   "public",
		},
	}
}
