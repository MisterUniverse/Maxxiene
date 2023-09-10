/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package notes

import (
	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var NotesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A place for taking notes",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
