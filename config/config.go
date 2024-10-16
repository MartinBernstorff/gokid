package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type GokidConfig struct {
	Trunk string
}

func Init() GokidConfig {
	configFile := findConfig(".gokid")
	if configFile != "" {
		viper.SetConfigFile(configFile)
		fmt.Println("Using config file:", configFile)
	} else {
		fmt.Println("No config file found")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}

	return GokidConfig{
		Trunk: viper.GetString("trunk"),
	}
}

func findConfig(configName string) string {
	configExtensions := []string{".yaml", ".yml", ".json", ".toml"}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}

	for {
		for _, ext := range configExtensions {
			configPath := filepath.Join(dir, configName+ext)
			fmt.Println("Exploring ", configPath)
			if _, err := os.Stat(configPath); err == nil {
				return configPath
			}
		}

		// Move to the parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root directory
			break
		}
		dir = parent
	}

	return ""
}
