package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DBConn string `mapstructure:"DB_CONN"`
	Port   string `mapstructure:"PORT"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetEnvPrefix("") // empty prefix to match exact env var names
	viper.AutomaticEnv()

	// Explicitly bind environment variables to struct fields
	viper.BindEnv("DB_CONN")
	viper.BindEnv("PORT")

	if _, err := os.Stat(".env"); err == nil {
		absPath, err := filepath.Abs(".env")
		if err != nil {
			log.Printf("Warning: could not get absolute path for .env: %v", err)
		} else {
			viper.SetConfigFile(absPath)
		}

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: error reading .env file: %v", err)
		} else {
			log.Printf("Loaded .env from: %s", absPath)
		}
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	fmt.Printf("Config loaded: %+v\n", config)

	return config, nil
}
