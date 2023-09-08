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

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// procCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// procCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
