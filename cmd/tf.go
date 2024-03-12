/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
			fmt.Println(err)
		}
		fmt.Printf("%v\n", conf)
		modules, err := terraform.GetModulesFromGitlab(conf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", len(modules))
		localModules, err := terraform.GetLocalModules()
		if err != nil {
			fmt.Println(err)
		}
		mergedModules := terraform.MergeModules(modules, *localModules)
		fmt.Printf("%v\n", mergedModules)

	},
}

func init() {
	rootCmd.AddCommand(tfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
