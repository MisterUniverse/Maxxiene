/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package sites

import (
	"github.com/spf13/cobra"
)

// sitesCmd represents the sites command
var SitesCmd = &cobra.Command{
	Use:   "sites",
	Short: "A simple bookmark manager",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
