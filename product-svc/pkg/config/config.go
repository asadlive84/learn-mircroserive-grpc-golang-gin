package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	PORT              string `mapstructure:"PORT"`
	AUTH_PORT         string `mapstructure:"AUTH_PORT"`
	JWT_SECRET_KEY    string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
