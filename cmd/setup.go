/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
// cmd/setup.go

package cmd

import (
	"fmt"
	mdb "maxx/pkg/db"
	"maxx/pkg/filemgr"
	"os"

	"github.com/spf13/cobra"
)

var workingDir string

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up initial configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Create directories
		directories := []string{workingDir, workingDir + "\\backups", workingDir + "\\config", workingDir + "\\data"}
		for _, dir := range directories {
			if err := filemgr.CreateDirectory(dir); err != nil {
				fmt.Println(err)
				return
			}
		}

		// Create .env file
		envValues := map[string]string{
			"TODO_FILE_PATH": workingDir + "\\todo.md",
			"CONFIG_DIR":     workingDir + "\\config",
			"DATA_DIR":       workingDir + "\\data",
			"BOOKMARKS":      workingDir + "\\data\\bookmarks.json",
			"BACKUPS":        workingDir + "\\backups",
			"DATABASE":       workingDir + "\\data\\maxxdb.db",
		}

		if err := filemgr.WriteEnvFile(envValues["CONFIG_DIR"]+"\\.env", envValues); err != nil {
			fmt.Println(err)
			return
		}
		mdb.MaxxDB = mdb.MaxxDataBase{}
		mdb.MaxxDB.Storage = mdb.NewDataStorage(envValues["DATA_DIR"] + "\\maxxdb.db")
		mdb.MaxxDB.Storage.InitializeTables()

		// application files
		fMap := map[string]string{
			"todo":      workingDir + "\\todo.md",
			"bookmarks": workingDir + "\\data\\bookmarks.json",
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
	workingDir = os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	if workingDir == "" {
		fmt.Println("Could not find LOCALAPPDATA environment variable")
		return
	}

	rootCmd.AddCommand(setupCmd)
}
