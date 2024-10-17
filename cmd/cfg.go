/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"gokid/config"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cfgCmd(write bool) {
	cfg := config.Init()

	if !write {
		prettyJSON, _ := json.MarshalIndent(cfg, "", "  ")
		fmt.Println(string(prettyJSON))
		return
	}

	viper.SafeWriteConfigAs(".gokid.yml")

	// Append gokid to gitignore if it exists
	gitignorePath := ".gitignore"
	if _, err := os.Stat(gitignorePath); err == nil {
		file, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening .gitignore file:", err)
			return
		}
		defer file.Close()

		contents, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading .gitignore file:", err)
			return
		}

		// Only append if the .gokid string is not already in the file
		ignoreString := ".gokid.*"
		if !strings.Contains(string(contents), ignoreString) {
			_, err = file.WriteString("\n" + ignoreString)
			if err != nil {
				fmt.Println("Error writing to .gitignore file:", err)
				return
			}
		}
	}
}

var newCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Print the identified cfg",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		write, _ := cmd.Flags().GetBool("write")
		cfgCmd(write)
	},
	Aliases: []string{"draft"},
}

func init() {
	newCmd.PersistentFlags().Bool("write", false, "Write the identified cfg")
	rootCmd.AddCommand(newCmd)
}
