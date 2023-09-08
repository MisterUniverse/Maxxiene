/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}
	TaskCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		todoFilePath := viper.GetString("TODO_FILE_PATH")
		if todoFilePath == "" {
			fmt.Println("TODO_FILE_PATH not set in .env, using default path: todo.md")
			todoFilePath = "todo.md"
		}
		addTaskToMarkdown(task, todoFilePath)
		fmt.Printf("Added a new task: \"%s\"\n", task)
	},
}

func addTaskToMarkdown(task, path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create file: %s\n", err)
		return
	}
	defer file.Close()

	formattedTask := fmt.Sprintf("- [ ] %s\n", task)

	if _, err := file.WriteString(formattedTask); err != nil {
		fmt.Printf("Failed to write to file: %s\n", err)
		return
	}
}
