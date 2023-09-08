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
	Short: "Add a bookmark",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := args[1]
		category := args[2]

		bookmarks = append(bookmarks, Bookmark{Name: name, URL: url, Category: category})
		bookmarkJSON, _ := json.Marshal(bookmarks)

		err := os.WriteFile(viper.GetString("BOOKMARKS"), bookmarkJSON, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Added bookmark: %s\n", name)
	},
}

func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData+"\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	addCmd.Args = cobra.ExactArgs(3)
	SitesCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
