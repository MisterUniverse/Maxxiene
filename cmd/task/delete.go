/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"maxx/pkg/db"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		db.MaxxDB.Storage = db.NewDataStorage(viper.GetString("DATABASE"))
		if len(args) == 0 {
			fmt.Println("Please provide a task id to delete.")
			return
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task number:", err)
			return
		}

		db.MaxxDB.Storage.DeleteData("todos", "id = ?", taskID)
	},
}

func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData + "\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	TaskCmd.AddCommand(deleteCmd)
}
