/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/volkovartem/joven/config"
	"github.com/volkovartem/joven/terraform"
)

// tfCmd represents the tf command
var tfCmd = &cobra.Command{
	Use:   "tf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Load()
		if err != nil {
			log.Fatalf(err.Error())
		}

		withMarkedOutdated, err := terraform.CompareGitLabModules(conf)
		if err != nil {
			log.Fatalf(err.Error())
		}
		terraform.Print(os.Stdout, withMarkedOutdated)

	},
}

func init() {
	rootCmd.AddCommand(tfCmd)
}
