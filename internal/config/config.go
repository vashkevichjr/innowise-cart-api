package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string
	PGDSN    string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	httpPort := viper.GetString("HTTP_PORT")
	if httpPort == "" {
		return nil, fmt.Errorf("environment variable HTTP_PORT is not set")
	}

	pgDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSL_MODE"))

	return &Config{HTTPPort: httpPort, PGDSN: pgDSN}, nil

}
