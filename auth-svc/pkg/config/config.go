package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseUserName string `mapstructure:"DATABASE_USER_NAME"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DbPassword       string `mapstructure:"DB_PASSWORD"`
	DbPort           string `mapstructure:"DB_PORT"`
	DbHost           string `mapstructure:"DB_HOST"`
	Port             string `mapstructure:"PORT"`
	DBUrl            string `mapstructure:"DB_URL"`
	JWTSecretKey     string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
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
