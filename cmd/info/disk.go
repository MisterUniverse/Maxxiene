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

var outdisk = `Get-Volume`

//var pipeout = `| Out-File ` // <- dont forget to add the path that's why the space is there

var diskUsageCmd = &cobra.Command{
	Use:   "disk",
	Short: "Prints disk usage of the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		posh := utils.NewShell()

		printResult := fmt.Sprintf("%s\n", outdisk)

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
	InfoCmd.AddCommand(diskUsageCmd)
}
