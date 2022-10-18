/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"os"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// browserCmd represents the browser command
var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Opens the default web browser",
	Long:  `Opens the default web browser along with any tabs that are set on the "tab.txt" located in the resources dir`,
	Run: func(cmd *cobra.Command, args []string) {
		tabs, _ := os.Open("./resources/tabs.txt")
		scanner := bufio.NewScanner(tabs)
		for scanner.Scan() {
			line := scanner.Text()
			browser.OpenURL(line)
		}
	},
}

func init() {
	rootCmd.AddCommand(browserCmd)
}
