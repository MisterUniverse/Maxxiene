/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// startupCmd represents the startup command
var startupCmd = &cobra.Command{
	Use:   "startup",
	Short: "Maxxiene is awake a ready to start working",
	Long:  `Here Maxx will will begin the morning routine and open all the browser windows we want to program her to open. along with other utilities we need to start working.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := "Desmond"
		fmt.Printf("Good morning %v, I am getting your tools ready. I will start by preparing a work environment and then opening some browser tabs for you. \n", user)

		envExist, err := fileOrDirExists("MaxxConfig.json")
		check(err)
		if !envExist {
			createWorkflowDirectories()
		}

		f, err := os.Open("./urls/startup_links.txt")
		check(err)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			browser.OpenURL(line)
		}
	},
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func createWorkflowDirectories() {
	workdirs := []string{"./Projects", "./Notes", "./Todos"}

	createConfig()

	for _, v := range workdirs {
		makeDir(v)
	}

	setupTodos()
	appendData()
}

func createConfig() {
	config, err := os.Create("MaxxConfig.json")
	check(err)
	defer config.Close()
}

func setupTodos() {
	path := "./Todos/main_todo_list.md"
	todo, err := os.Create(path)
	check(err)
	defer todo.Close()
}
func appendData() {
	content := "# Todo List"
	path := "./Todos/main_todo_list.md"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	check(err)
	file.Write([]byte(content))
}
func makeDir(dirName string) {
	exist, err := fileOrDirExists(dirName)
	check(err)
	if !exist {
		err := os.Mkdir(dirName, 0755)
		check(err)
	}
}

func fileOrDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getWorkingDir() {
	dir, err := os.Getwd()
	check(err)
	fmt.Println(dir)
}

func getDirContents(path string) {
	/*entries, err := os.ReadDir(path)
	check(err)
	//iterate through dir and print name
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}*/
}

func init() {
	rootCmd.AddCommand(startupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
