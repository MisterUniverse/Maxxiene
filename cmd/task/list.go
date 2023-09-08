/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		todoFilePath := viper.GetString("TODO_FILE_PATH")
		if todoFilePath == "" {
			fmt.Println("TODO_FILE_PATH not set in .env, using default path: todo.md")
			todoFilePath = "todo.md"
		}
		listTasksFromMarkdown(todoFilePath)
	},
}

func listTasksFromMarkdown(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Task List:")

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "- [ ] ") || strings.HasPrefix(line, "- [x] ") {
			task := strings.TrimPrefix(line, "- [ ] ")
			task = strings.TrimPrefix(task, "- [x] ")
			fmt.Printf("%d. %s\n", i, task)
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
}
func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData+"\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	TaskCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
