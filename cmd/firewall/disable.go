/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package firewall

import (
	"maxx/pkg/winfirewall"

	"github.com/spf13/cobra"
)

// var all, current, domain, global, private bool

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable firewall for a specified profile",
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
		err := winfirewall.DisableFirewall(selected)
		cobra.CheckErr(err)
	},
}

func init() {
	disableCmd.Flags().BoolVarP(&all, "all", "a", false, "enables firewall for allprofiles")
	disableCmd.Flags().BoolVarP(&current, "current", "c", false, "enables firewall for currentprofile")
	disableCmd.Flags().BoolVarP(&domain, "domain", "d", false, "enables firewall for domainprofile")
	disableCmd.Flags().BoolVarP(&global, "global", "g", false, "enables firewall for global")
	disableCmd.Flags().BoolVarP(&private, "private", "p", false, "enables firewall for privateprofile")
	FirewallCmd.AddCommand(disableCmd)
}
