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

func init() {}
