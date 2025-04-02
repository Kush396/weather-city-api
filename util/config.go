package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
