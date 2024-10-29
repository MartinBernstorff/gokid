package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"

	"github.com/spf13/viper"
)

type GokidConfig struct {
	AutoMerge       bool
	BranchPrefix    string
	BranchSuffix    string
	Draft           bool
	ForceMerge      bool
	MergeStrategy   string
	PreMergeCommand string
	Trunk           string
}

func NewConfig(autoMerge bool, branchPrefix string, branchSuffix string, draft bool, forceMerge bool, mergeStrategy string, preMergeCommand string, trunk string) GokidConfig {
	validateMergeStrategy(mergeStrategy)
	validateForceMerge(forceMerge, preMergeCommand, autoMerge)

	return GokidConfig{
		AutoMerge:       autoMerge,
		BranchPrefix:    branchPrefix,
		BranchSuffix:    branchSuffix,
		Draft:           draft,
		ForceMerge:      forceMerge,
		MergeStrategy:   mergeStrategy,
		PreMergeCommand: preMergeCommand,
		Trunk:           trunk,
	}
}

func Defaults() GokidConfig {
	return NewConfig(false, "", "", false, false, "merge", "", "main")
}

func Load(configName string) GokidConfig {
	defaults := Defaults()
	defaultValue := reflect.ValueOf(defaults)

	// Set viper defaults based on struct fields
	for _, field := range reflect.VisibleFields(reflect.TypeOf(defaults)) {
		viper.SetDefault(field.Name, defaultValue.FieldByName(field.Name).Interface())
	}

	configFile := findConfig(configName)
	if configFile != "" {
		viper.SetConfigFile(configFile)
		file, err := os.Open(configFile)
		if err != nil {
			fmt.Println("Error opening config file:", err)
			return GokidConfig{}
		}
		defer file.Close()

		readErr := viper.ReadConfig(file)
		if readErr != nil {
			fmt.Println("Error reading config file:", readErr)
			panic(readErr)
		}

		fmt.Println("Using config file:", configFile)
	} else {
		fmt.Println("No config file found")
	}

	return NewConfig(
		viper.GetBool("automerge"),
		viper.GetString("branchprefix"),
		viper.GetString("branchsuffix"),
		viper.GetBool("draft"),
		viper.GetBool("forcemerge"),
		viper.GetString("mergestrategy"),
		viper.GetString("premergecommand"),
		viper.GetString("trunk"),
	)
}

func validateMergeStrategy(mergeStrategy string) {
	allowedStrategies := []string{"squash", "rebase", "merge"}

	if !slices.Contains(allowedStrategies, mergeStrategy) {
		msg := fmt.Sprintf("Merge strategy %s not allowed, allowed are: %s", mergeStrategy, strings.Join(allowedStrategies, ", "))
		panic(msg)
	}
}

func validateForceMerge(forceMerge bool, preMergeCommand string, autoMerge bool) {
	if forceMerge && preMergeCommand == "" {
		panic("Force merge can only be enabled when pre-merge command is set")
	}

	if autoMerge && forceMerge {
		panic("Either auto merge or force merge can be enabled, not both")
	}
}

func findConfig(configName string) string {
	configExtensions := []string{".yaml", ".yml", ".json", ".toml"}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}

	// Look in all parent directories for the config file
	for {
		for _, ext := range configExtensions {
			configPath := filepath.Join(dir, configName+ext)
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
