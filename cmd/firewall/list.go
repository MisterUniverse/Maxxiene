/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"fmt"
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

var ruleName string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all firewall rules",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := winfirewall.ListFirewallRules(ruleName)
		cobra.CheckErr(err)
		fmt.Println(string(output), err)
	},
}

func init() {
	listCmd.Flags().StringVarP(&ruleName, "name", "n", "all", "search by name or leave empty to search all rules")
	FirewallCmd.AddCommand(listCmd)
}
