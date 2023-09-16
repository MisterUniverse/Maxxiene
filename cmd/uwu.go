/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"maxx/pkg/waifuim"

	"github.com/spf13/cobra"
)

var meow bool

// uwuCmd represents the uwu command
var uwuCmd = &cobra.Command{
	Use:   "uwu",
	Short: "Summon a random waifu",
	Run: func(cmd *cobra.Command, args []string) {
		switch meow {
		case true:
			waifuim.GetWaifuMeow()
		default:
			waifuim.GetWaifu()
		}
	},
}

func init() {
	uwuCmd.Flags().BoolVar(&meow, "meow", false, "uwu meow...")
	rootCmd.AddCommand(uwuCmd)
}
