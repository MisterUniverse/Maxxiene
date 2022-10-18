/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list of things to do",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("========================")
		b := fm.ReadFile("./todo.md")
		fmt.Println("\n" + string(b))
		fmt.Println("========================")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
