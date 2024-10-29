/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"gokid/config"

	"github.com/spf13/cobra"
)

var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Print the identified config",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		showDefaults, _ := cmd.Flags().GetBool("defaults")

		var cfg config.GokidConfig
		if showDefaults {
			cfg = config.Defaults()
		} else {
			cfg = config.Load(config.DefaultFileName)
		}

		prettyJSON, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling config:", err)
			return
		}

		fmt.Println(string(prettyJSON))
	},
	Aliases: []string{"c"},
}

func init() {
	cfgCmd.Flags().Bool("defaults", false, "Show default configuration values")
	rootCmd.AddCommand(cfgCmd)
}
