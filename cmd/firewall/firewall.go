/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"github.com/spf13/cobra"
)

// FirewallCmd represents the firewall command
var FirewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Modify firewall rules",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
