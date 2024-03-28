/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/netvolart/joven/cdk"
	"github.com/netvolart/joven/config"
	"github.com/spf13/cobra"
)

var cdkCmd = &cobra.Command{
	Use:   "cdk",
	Short: "Prints outdated Terraform modules form a local project",
	Long: `This command compares the versions of the Terraform modules in the 
	local project with the latest versions available in the GitLab registry 
	and Community terraform registry. Require to run terraform get before.`,
	Run: func(cmd *cobra.Command, args []string) {

		_, err := config.Load()

		if err != nil {
			log.Fatalf(err.Error())
		}

		cdk.CompareCDKConstructs()

		// withMarkedOutdated, err := terraform.CompareGitLabModules(conf, localModulesData)
		// if err != nil {
		// 	log.Fatalf(err.Error())
		// }
		// terraform.Print(os.Stdout, withMarkedOutdated)

	},
}

func init() {
	rootCmd.AddCommand(cdkCmd)

}
