/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"maxx/utils"
	"strings"

	"github.com/spf13/cobra"
)

// ipconfCmd represents the ipconf command
var ipconfCmd = &cobra.Command{
	Use:   "ipconf",
	Short: "Shows the current ip configuration.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pshCmd := "netsh interface ipv4 show config"
		psh := utils.NewShell()

		stdOut, stdErr, err := psh.Execute(pshCmd)
		fmt.Printf("\n ipconf: \n StdOut : '%s' \n StdErr: '%s' \n Err: %s", strings.TrimSpace(stdOut), stdErr, err)
	},
}

func init() {
	NetCmd.AddCommand(ipconfCmd)
}
