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

var ram = `Get-WmiObject win32_physicalmemory | Format-Table Manufacturer,Banklabel,Configuredclockspeed,Devicelocator,Capacity,Serialnumber -autosize`
var cpu = `Get-WmiObject Win32_Processor`

//var pipeout = `| Out-File ` // <- dont forget to add the path that's why the space is there

var ramcpuCmd = &cobra.Command{
	Use:   "hardware",
	Short: "Gets hardware information",
	Long:  `Uses powershell to grab the ram and cpu information and prints it to the screen`,
	Run: func(cmd *cobra.Command, args []string) {
		posh := utils.NewShell()

		printResult := fmt.Sprintf("%s\n %s\n", ram, cpu)

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
	InfoCmd.AddCommand(ramcpuCmd)
}
