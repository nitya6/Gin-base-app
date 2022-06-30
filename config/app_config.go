package config

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

// Configurations exported
type AppConfigurations struct {
	Server   ServerConfigurations   `mapstructure:"server"`
	Database DatabaseConfigurations `mapstructure:"database"`
	Profile  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port          string
	HMACKey       string `mapstructure:"hmacKey"`
	MetricsPath   string `mapstructure:"metricsPath"`
	MetricsPort   string `mapstructure:"metricsPort"`
	MetricsPrefix string `mapstructure:"metricsPrefix"`
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	Uri string `mapstructure:"uri"`
}

var AppConfig AppConfigurations

func LoadAppConfig() error {

	viper.SetConfigName("config")

	absPath, _ := filepath.Abs("./config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(absPath)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.BindEnv("database.dbuser", "DB_USER")
	viper.BindEnv("database.dbpass", "DB_PASS")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return errors.New("failed to loading config")

	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return errors.New("failed to loading config")

	}
	return nil
}
