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

func init() {}
