package config

import (
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Auth struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"auth"`
}

// LoadConfig reads configuration from a file and initializes the config struct
func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName(env)     // Name of config file (without extension)
	viper.AddConfigPath("./env") // Path to look for the config file in
	viper.SetConfigType("toml")  // Type of config file

	// If the configuration file is missing, Viper will return an error
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	// Unmarshal the configuration values into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
