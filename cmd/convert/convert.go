/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package convert

import (
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A palette for converting files",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
