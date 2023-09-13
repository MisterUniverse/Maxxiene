/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

var delName string

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a firewall rule",
	Run: func(cmd *cobra.Command, args []string) {
		err := winfirewall.DeleteFirewallRule(delName)
		cobra.CheckErr(err)
	},
}

func init() {
	delCmd.Flags().StringVarP(&delName, "name", "n", "", "name of rule to delete")
	FirewallCmd.AddCommand(delCmd)
}
