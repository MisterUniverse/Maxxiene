/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

var exportName string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export [path]",
	Short: "Export firewall rules binary",
	Run: func(cmd *cobra.Command, args []string) {
		winfirewall.ExportFirewallRules(exportName)
		winfirewall.ExportHumanReadable(exportName)
	},
}

func init() {
	exportCmd.Flags().StringVarP(&exportName, "name", "n", "firewall_rules", "name of file")
	FirewallCmd.AddCommand(exportCmd)
}
