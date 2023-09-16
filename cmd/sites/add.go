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

		bookmarks = append(bookmarks, Bookmark{Name: name, URL: url, Category: category})
		bookmarkJSON, _ := json.Marshal(bookmarks)

		err := os.WriteFile(viper.GetString("paths.BOOKMARKS"), bookmarkJSON, 0644)
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
