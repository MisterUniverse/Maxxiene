/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package convert

import (
	"fmt"
	"os"

	"github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)

// md2htmlCmd represents the md2html command
var md2htmlCmd = &cobra.Command{
	Use:   "md2html",
	Short: "Convert a Markdown file to HTML",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must specify a Markdown file to convert")
			return
		}

		mdPath := args[0]
		mdContent, err := os.ReadFile(mdPath)
		if err != nil {
			fmt.Printf("Could not read file: %v\n", err)
			return
		}

		htmlContent := blackfriday.Run(mdContent)
		htmlPath := mdPath + ".html"

		err = os.WriteFile(htmlPath, htmlContent, 0644)
		if err != nil {
			fmt.Printf("Could not write to file: %v\n", err)
			return
		}

		fmt.Printf("Successfully converted %s to %s\n", mdPath, htmlPath)
	},
}

func init() {
	ConvertCmd.AddCommand(md2htmlCmd)
}
