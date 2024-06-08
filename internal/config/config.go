package config

type Config struct {
	DB *DBConfig
}

func LoadConfig() *Config {
	return &Config{
		DB: loadDatabaseConfig(),
	}
}
