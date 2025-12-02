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

type envConfig struct {
	HTTPPort   string `mapstructure:"HTTP_PORT" validate:"required"`
	dbUser     string `mapstructure:"DB_USER" validate:"required"`
	dbPassword string `mapstructure:"DB_PASSWORD" validate:"required"`
	dbName     string `mapstructure:"DB_NAME" validate:"required"`
	dbHost     string `mapstructure:"DB_HOST" validate:"required"`
	dbPort     string `mapstructure:"DB_PORT" validate:"required"`
	sslMode    string `mapstructure:"SSL_MODE" validate:"required,oneof=disable"`
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

	cfg := envConfig{
		HTTPPort:   viper.GetString("HTTP_PORT"),
		dbUser:     viper.GetString("DB_USER"),
		dbPassword: viper.GetString("DB_PASSWORD"),
		dbName:     viper.GetString("DB_NAME"),
		dbHost:     viper.GetString("DB_HOST"),
		dbPort:     viper.GetString("DB_PORT"),
		sslMode:    viper.GetString("SSL_MODE"),
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	pgDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.dbUser,
		cfg.dbPassword,
		cfg.dbHost,
		cfg.dbPort,
		cfg.dbName,
		cfg.sslMode)

	return &Config{HTTPPort: cfg.dbPort, PGDSN: pgDSN}, nil

}
