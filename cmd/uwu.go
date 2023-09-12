/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"maxx/pkg/waifuim"

	"github.com/spf13/cobra"
)

// uwuCmd represents the uwu command
var uwuCmd = &cobra.Command{
	Use:   "uwu",
	Short: "random waifu",
	Run: func(cmd *cobra.Command, args []string) {
		waifuim.GetWaifu()
	},
}

func init() {
	rootCmd.AddCommand(uwuCmd)
}
