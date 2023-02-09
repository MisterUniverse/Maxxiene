/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package env

import (
	"fmt"
	"maxx/utils"
	"strings"

	"github.com/spf13/cobra"
)

var (
	exe = `Invoke-Expression -Command C:\\don't-hard-code-path\\windowsenv.ps1`
)
var initWinCmd = &cobra.Command{
	Use:   "initWin",
	Short: "A brief description of your command",
	Long:  `Initializes windows env`,
	Run: func(cmd *cobra.Command, args []string) {
		posh := utils.NewShell()

		stdOut, stdErr, err := posh.Execute(exe)
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
	EnvCmd.AddCommand(initWinCmd)
}
