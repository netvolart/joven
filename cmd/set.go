/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	conf "github.com/volkovartem/joven/config"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		groups, _ := cmd.Flags().GetString("groups")
		groupsList := strings.Split(groups, ",")
		configData := conf.New(token, groupsList)
		err := configData.Save()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("token", "t", "", "Gitlab personal access token")
	setCmd.MarkFlagRequired("token")
	setCmd.Flags().StringP("groups", "g", "", "Comma separated list of GitLab groups with Terraform modules and CDK libs")
	setCmd.MarkFlagRequired("groups")
}
