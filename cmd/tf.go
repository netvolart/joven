/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/netvolart/joven/config"
	"github.com/netvolart/joven/terraform"
	"github.com/spf13/cobra"
)

// tfCmd represents the tf command
var tfCmd = &cobra.Command{
	Use:   "tf",
	Short: "Prints outdated Terraform modules form a local project",
	Long: `This command compares the versions of the Terraform modules in the 
	local project with the latest versions available in the GitLab registry 
	and Community terraform registry. Require to run terraform get before.`,
	Run: func(cmd *cobra.Command, args []string) {

		conf, err := config.Load()

		if err != nil {
			log.Fatalf(err.Error())
		}
		localModulesData, err := os.ReadFile(".terraform/modules/modules.json")

		if err != nil {
			log.Fatalf(err.Error())
		}
		withMarkedOutdated, err := terraform.CompareGitLabModules(conf, localModulesData)
		if err != nil {
			log.Fatalf(err.Error())
		}
		terraform.Print(os.Stdout, withMarkedOutdated)

	},
}

func init() {
	rootCmd.AddCommand(tfCmd)

}
