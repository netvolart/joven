/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/netvolart/joven/internal/cdk"
	"github.com/netvolart/joven/internal/config"
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

		withMarkedOutdated, err := cdk.CompareCDKConstructs()
		if err != nil {
			log.Fatalf(err.Error())
		}
		cdk.Print(os.Stdout, withMarkedOutdated)

	},
}

func init() {
	rootCmd.AddCommand(cdkCmd)

}
