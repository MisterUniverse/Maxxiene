/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"os"
	"maxx/pkg/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db.MaxxDB.Storage = db.NewDataStorage(viper.GetString("DATABASE"))

		var items []db.ItemScanner
		var err error

		items, err = db.MaxxDB.Storage.ListItems("todos", &db.Todo{})
		logError(err)

		fmt.Println(len(items))

		for _, item := range items {
			switch v := item.(type) {
			case *db.Todo:
				fmt.Printf("ID: %d, Completed: %v, Task: %v\n", v.ID, v.Completed, v.Task)
			}
		}
	},
}

func logError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData+"\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	TaskCmd.AddCommand(listCmd)
}
