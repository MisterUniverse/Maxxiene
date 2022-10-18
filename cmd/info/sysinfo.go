/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"
	"maxx/utils"
	"strings"

	"github.com/spf13/cobra"
)

var sysinfoCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sysInfo := `Get-ComputerInfo`
		//pipeout := `| Out-File ` // <- dont forget to add the path that's why the space is there

		posh := utils.NewShell()

		printResult := fmt.Sprintf("%s\n", sysInfo)

		stdOut, stdErr, err := posh.Execute(printResult)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf(
			"\nLog System Information:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s",
			strings.TrimSpace(stdOut),
			stdErr,
			err,
		)
	},
}

func init() {
	InfoCmd.AddCommand(sysinfoCmd)
}
