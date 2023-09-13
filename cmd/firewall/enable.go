/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

var all, current, domain, global, private bool

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable firewall for a specified profile",
	Run: func(cmd *cobra.Command, args []string) {
		var selected string
		switch {
		case all:
			selected = "allprofiles"
		case current:
			selected = "currentprofile"
		case domain:
			selected = "domainprofile"
		case global:
			selected = "global"
		case private:
			selected = "privateprofile"
		default:
			selected = "publicprofile"
		}
		err := winfirewall.EnableFirewall(selected)
		cobra.CheckErr(err)
	},
}

func init() {
	enableCmd.Flags().BoolVarP(&all, "all", "a", false, "enables firewall for allprofiles")
	enableCmd.Flags().BoolVarP(&current, "current", "c", false, "enables firewall for currentprofile")
	enableCmd.Flags().BoolVarP(&domain, "domain", "d", false, "enables firewall for domainprofile")
	enableCmd.Flags().BoolVarP(&global, "global", "g", false, "enables firewall for global")
	enableCmd.Flags().BoolVarP(&private, "private", "p", false, "enables firewall for privateprofile")
	FirewallCmd.AddCommand(enableCmd)
}
