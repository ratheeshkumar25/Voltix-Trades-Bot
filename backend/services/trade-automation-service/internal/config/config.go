package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	NatsURL string `mapstructure:"NATS_URL"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.Port == "" {
		config.Port = "3002"
	}
	if config.NatsURL == "" {
		config.NatsURL = "nats://localhost:4222"
	}

	return &config, nil
}
