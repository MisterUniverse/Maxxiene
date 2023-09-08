/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"maxx/pkg/filemgr"

	"github.com/spf13/cobra"
)

// shredCmd represents the shred command
var shredCmd = &cobra.Command{
	Use:   "shred",
	Short: "Shred a file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must specify a directory to shred.")
			return
		}

		directoryPath := args[0]
		if filemgr.IsDir(directoryPath) {
			err := filemgr.ShredDirectory(directoryPath, 3) // overwrite 3 times
			if err != nil {
				fmt.Printf("shred failed: %v\n", err)
			}
			return
		}

		err := filemgr.ShredFile(directoryPath, 3)
		if err != nil {
			fmt.Printf("shred failed: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(shredCmd)
}
