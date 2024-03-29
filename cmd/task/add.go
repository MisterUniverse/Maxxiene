/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"maxx/pkg/db"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Run: func(cmd *cobra.Command, args []string) {
		db.MaxxDB.Storage = db.NewDataStorage(viper.GetString("paths.DATABASE"))
		joined := strings.Join(args, " ")

		db.MaxxDB.Storage.InsertData("todos", "completed, task", false, joined)
		fmt.Printf("Added a new task: \"%s\"\n", joined)
	},
}

func init() {
	TaskCmd.AddCommand(addCmd)
}
