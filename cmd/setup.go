/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
// cmd/setup.go

package cmd

import (
	"fmt"

	"maxx/pkg/filemgr"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up initial configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Create directories
		directories := []string{"backups", "config", "data"}
		for _, dir := range directories {
			if err := filemgr.CreateDirectory(dir); err != nil {
				fmt.Println(err)
				return
			}
		}

		// Create .env file
		envValues := map[string]string{
			"TODO_FILE_PATH": "./todo.md",
			"CONFIG_DIR":     "./config",
			"DATA_DIR":       "./data",
			"BOOKMARKS":      "./data/bookmarks.json",
			"BACKUPS":        "./backups",
		}

		if err := filemgr.WriteEnvFile(envValues["CONFIG_DIR"]+"/.env", envValues); err != nil {
			fmt.Println(err)
			return
		}

		// application files
		fMap := map[string]string{
			"todo":      "./todo.md",
			"bookmarks": "./data/bookmarks.json",
		}

		// Create additional files
		files := []string{fMap["todo"], fMap["bookmarks"]}
		for _, file := range files {
			if err := filemgr.CreateFile(file); err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("Setup complete.")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
