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

var port string

// rppCmd represents the rpp command
var rppCmd = &cobra.Command{
	Use:   "rpp",
	Short: "running process port",
	Long: `The running process port (rpp) command runs a powershell command 
that checks to see which process is running on a specific port.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(port) == 0 {
			cmd.Help()
			return
		}

		pshCmd := fmt.Sprintf("Get-Process -Id (Get-NetTCPConnection -LocalPort %v).OwningProcess", port)
		psh := utils.NewShell()

		stdOut, stdErr, err := psh.Execute(pshCmd)
		fmt.Printf("\n rpp: \n StdOut : '%s' \n StdErr: '%s' \n Err: %s", strings.TrimSpace(stdOut), stdErr, err)
	},
}

func init() {
	rppCmd.Flags().StringVarP(&port, "port", "p", "", "Port to ping")

	if err := rppCmd.MarkFlagRequired("port"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(rppCmd)
}
