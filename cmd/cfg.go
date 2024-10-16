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

func cfg() {
	cfg := config.Init()
	prettyJSON, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(prettyJSON))
}

// draftCmd represents the draft command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Print the identified cfg",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfg()
	},
	Aliases: []string{"draft"},
}

func init() {
	rootCmd.AddCommand(cfgCmd)
}
