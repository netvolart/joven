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
	Short: "Prints outdated CDK dependencies in the project.",

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
