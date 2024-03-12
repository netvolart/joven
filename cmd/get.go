/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	conf "github.com/volkovartem/joven/config"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := conf.Load()
		if err != nil {
			fmt.Println(err)
		}
		json, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(json))
		}
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
