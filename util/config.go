package util

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	VesApiUrl         string `mapstructure:"VES_API_URL"`
	VesApiKey         string `mapstructure:"VES_API_KEY"`
	AllowedOriginsStr string `mapstructure:"ALLOWED_ORIGINS"`
	AllowedOrigins    []string
}

func LoadConfig(path string) (*Config, error) {
	// viper.AddConfigPath(path)
	// viper.SetConfigType("env")

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.AllowedOrigins = strings.Split(config.AllowedOriginsStr, ",")

	return &config, nil
}
