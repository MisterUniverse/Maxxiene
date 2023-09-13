/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

var fileName string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a firewall rule binary",
	Run: func(cmd *cobra.Command, args []string) {
		err := winfirewall.ImportFirewallRules(fileName)
		cobra.CheckErr(err)
	},
}

func init() {
	importCmd.Flags().StringVarP(&fileName, "path", "p", "", "path/to/file.bin")
	FirewallCmd.AddCommand(importCmd)
}
