/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a firewall rule",
	Run: func(cmd *cobra.Command, args []string) {
		err := winfirewall.AddFirewallRule()
		cobra.CheckErr(err)
	},
}

func init() {
	FirewallCmd.AddCommand(addCmd)
}
