/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	conf "github.com/volkovartem/joven/config"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration for the tool",
	Long: `You have to provide a Gitlab personal token with --token flag and 
	GitLab group that is a parrent to your Terraform moduels with --groups flag. 
	Config will be stored locally`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		groups, _ := cmd.Flags().GetString("groups")
		groupsList := strings.Split(groups, ",")
		configData := conf.New(token, groupsList)
		err := configData.Save()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("token", "t", "", "Gitlab personal access token")
	if err := setCmd.MarkFlagRequired("token"); err != nil {
		log.Println(err)
	}

	setCmd.Flags().StringP("groups", "g", "", "Comma separated list of GitLab groups with Terraform modules and CDK libs")
	if err := setCmd.MarkFlagRequired("groups"); err != nil {
		log.Println(err)
	}
}
