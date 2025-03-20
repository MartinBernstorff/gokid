package cmd

import (
	"fmt"
	"gokid/config"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// p3: This function can likely be refactored
func updateGitignore() error {
	gitignorePath := ".gitignore"
	ignoreString := ".gokid.*"

	// Create .gitignore if it doesn't exist
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		file, err := os.OpenFile(gitignorePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("error creating .gitignore file: %w", err)
		}
		defer file.Close()
		_, err = file.WriteString(ignoreString + "\n")
		return fmt.Errorf("error writing to .gitignore file: %w", err)
	}

	// Update existing .gitignore
	file, err := os.OpenFile(gitignorePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening .gitignore file: %w", err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading .gitignore file: %w", err)
	}

	if !strings.Contains(string(contents), ignoreString) || len(contents) == 0 {
		if !strings.HasSuffix(string(contents), "\n") {
			ignoreString = "\n" + ignoreString
		}
		_, err = file.WriteString(ignoreString + "\n")
		if err != nil {
			return fmt.Errorf("error writing to .gitignore file: %w", err)
		}
	}

	return nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gokid configuration",
	Long:  "Creates a default configuration file and updates .gitignore",
	Run: func(_ *cobra.Command, _ []string) {
		config.Load(config.DefaultFileName) // Load or create default config

		if err := viper.WriteConfigAs(config.DefaultFileName + ".yml"); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing config file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Created default configuration file:", config.DefaultFileName+".yml")

		if err := updateGitignore(); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating .gitignore: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Successfully initialized gokid configuration")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
