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
		cfg := config.Init()
		prettyJSON, _ := json.MarshalIndent(cfg, "", "  ")
		fmt.Println(string(prettyJSON))
	},
	Aliases: []string{"c"},
}

func init() {
	rootCmd.AddCommand(cfgCmd)
}
