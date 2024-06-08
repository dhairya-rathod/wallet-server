package config

import (
	"fmt"
	utils "wallet-server/internal/utils"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadDatabaseConfig() *DBConfig {
	return &DBConfig{
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnv("DB_PORT", "5432"),
		User:     utils.GetEnv("DB_USERNAME", "postgres"),
		Password: utils.GetEnv("DB_PASSWORD", "password"),
		DBName:   utils.GetEnv("DB_DATABASE", "go_api"),
		SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
	}
}

func (c *DBConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}
