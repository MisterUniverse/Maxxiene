/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"os"
	"strings"
	"maxx/pkg/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Run: func(cmd *cobra.Command, args []string) {
		db.MaxxDB.Storage = db.NewDataStorage(viper.GetString("DATABASE"))
		joined := strings.Join(args, " ")
		// trimmed := strings.Trim(joined, "[]")

		db.MaxxDB.Storage.InsertData("todos", "completed, task", false, joined)
		fmt.Printf("Added a new task: \"%s\"\n", joined)
	},
}

func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData+"\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}
	TaskCmd.AddCommand(addCmd)
}