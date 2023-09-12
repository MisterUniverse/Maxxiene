/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package proc

import (
	"github.com/spf13/cobra"
)

// procCmd represents the proc command
var ProcCmd = &cobra.Command{
	Use:   "proc",
	Short: "Process utility commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
