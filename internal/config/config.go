package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string `validate:"required,numeric"`
	PGDSN    string `validate:"required"`
}

type envConfig struct {
	HTTPPort string `mapstructure:"HTTP_PORT" validate:"required,numeric"`
	user     string `mapstructure:"DB_USER" validate:"required,oneof=disable"`
	password string `mapstructure:"DB_PASSWORD" validate:"required,oneof=disable"`
	name     string `mapstructure:"DB_NAME" validate:"required,oneof=disable"`
	host     string `mapstructure:"DB_HOST" validate:"required,oneof=disable"`
	port     string `mapstructure:"DB_PORT" validate:"required,oneof=disable"`
	sslmode  string `mapstructure:"DB_SSL_MODE" validate:"required,oneof=disable"`
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
