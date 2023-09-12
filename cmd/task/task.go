/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task/todo manager",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
