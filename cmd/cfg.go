/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gokid/config"

	"github.com/spf13/cobra"
)

func cfg() {
	cfg := config.Init()
	fmt.Print(cfg)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
