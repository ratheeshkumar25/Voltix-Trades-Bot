package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSL_MODE"`
}

func LoadConfig() (*Config, error) {
	//auto load env variables
	var config Config

	viper.AutomaticEnv()
	viper.SetConfigFile("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config file:", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("Unable to decode into struct:", err)
	}

	return &config, nil

}
