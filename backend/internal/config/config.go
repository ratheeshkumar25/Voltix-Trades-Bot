package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBName               string `mapstructure:"DB_NAME"`
	DBSSLMode            string `mapstructure:"DB_SSL_MODE"`
	VOLTIX_PORT          string `mapstructure:"VOLTIX_PORT"`
	GOOGLE_CLIENT_ID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GOOGLE_CLIENT_SECRET string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GOOGLE_REDIRECT_URL  string `mapstructure:"GOOGLE_REDIRECT_URL"`
}

func LoadConfig() (*Config, error) {
	//auto load env variables
	var config Config

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName("env")
	viper.AddConfigPath(".")

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
