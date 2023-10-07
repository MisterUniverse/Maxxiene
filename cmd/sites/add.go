/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package sites

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Bookmark struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Category string `json:"category"`
}

var bookmarks []Bookmark

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a bookmark {name} {url} {category}",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := args[1]
		category := args[2]

		// Step 1: Read the existing JSON file
		filePath := viper.GetString("paths.BOOKMARKS")
		fileContent, err := os.ReadFile(filePath)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("Error reading file:", err)
			return
		}

		// If file exists and has content, unmarshal it
		if len(fileContent) > 0 {
			err = json.Unmarshal(fileContent, &bookmarks)
			if err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				return
			}
		}

		// Step 2 & 3: Append new Bookmark to the slice
		bookmarks = append(bookmarks, Bookmark{Name: name, URL: url, Category: category})

		// Step 4: Marshal the updated slice back into JSON
		bookmarkJSON, err := json.Marshal(bookmarks)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// Step 5: Write the updated JSON back to the file
		err = os.WriteFile(filePath, bookmarkJSON, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Added bookmark: %s\n", name)
	},
}

func init() {
	addCmd.Args = cobra.ExactArgs(3)
	SitesCmd.AddCommand(addCmd)
}
