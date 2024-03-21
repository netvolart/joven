/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	conf "github.com/netvolart/joven/config"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List config file",
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
