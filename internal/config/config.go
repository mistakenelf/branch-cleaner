package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// SettingsConfig struct represents the config for the settings.
type SettingsConfig struct {
	EnableLogging     bool     `mapstructure:"enable_logging"`
	ProtectedBranches []string `mapstructure:"protected_branches"`
}

// Config represents the main config for the application.
type Config struct {
	Settings SettingsConfig `mapstructure:"settings"`
}

// LoadConfig loads a users config and creates the config if it does not exist
// located at ~/.branch-cleaner.yml.
func LoadConfig() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".branch-cleaner")
	viper.SetConfigType("yml")

	// Set default config values.
	viper.SetDefault("settings.enable_logging", false)
	viper.SetDefault("settings.protected_branches", []string{"main", "master", "develop", "dev", "prod"})

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}
}

// GetConfig returns the users config.
func GetConfig() (config Config) {
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error parsing config", err)
	}

	return
}
