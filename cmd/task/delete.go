/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task number to delete.")
			return
		}

		taskIndex, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task number:", err)
			return
		}

		todoFilePath := viper.GetString("TODO_FILE_PATH")
		if todoFilePath == "" {
			fmt.Println("TODO_FILE_PATH not set in .env, using default path: todo.md")
			todoFilePath = "todo.md"
		}

		deleteTaskFromMarkdown(taskIndex, todoFilePath)
	},
}

func deleteTaskFromMarkdown(taskIndex int, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	if taskIndex <= 0 || taskIndex > len(lines) {
		fmt.Println("Invalid task number")
		return
	}

	task := lines[taskIndex-1]
	lines = append(lines[:taskIndex-1], lines[taskIndex:]...)

	// Rewrite the file
	file, err = os.Create(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	fmt.Printf("Removed the task: \"%s\"\n", strings.TrimPrefix(task, "- [ ] "))
}

func init() {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	TaskCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
