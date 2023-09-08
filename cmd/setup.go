/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
// cmd/setup.go

package cmd

import (
	"fmt"
	"os"

	"maxx/pkg/filemgr"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up initial configuration",
	Run: func(cmd *cobra.Command, args []string) {
		localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
		if localAppData == "" {
			fmt.Println("Could not find LOCALAPPDATA environment variable")
			return
		}

		// Create directories
		directories := []string{localAppData, localAppData+"\\backups", localAppData+"\\config", localAppData+"\\data"}
		for _, dir := range directories {
			if err := filemgr.CreateDirectory(dir); err != nil {
				fmt.Println(err)
				return
			}
		}

		// Create .env file
		envValues := map[string]string{
			"TODO_FILE_PATH": localAppData+"\\todo.md",
			"CONFIG_DIR":     localAppData+"\\config",
			"DATA_DIR":       localAppData+"\\data",
			"BOOKMARKS":      localAppData+"\\data\\bookmarks.json",
			"BACKUPS":        localAppData+"\\backups",
		}

		if err := filemgr.WriteEnvFile(envValues["CONFIG_DIR"]+"\\.env", envValues); err != nil {
			fmt.Println(err)
			return
		}

		// application files
		fMap := map[string]string{
			"todo":      localAppData+"\\todo.md",
			"bookmarks": localAppData+"\\data\\bookmarks.json",
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
