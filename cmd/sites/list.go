/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package sites

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ByCategory []Bookmark

func (a ByCategory) Len() int      { return len(a) }
func (a ByCategory) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCategory) Less(i, j int) bool {
	return a[i].Category < a[j].Category
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all bookmarks",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(viper.GetString("BOOKMARKS"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		json.Unmarshal(data, &bookmarks)

		// Sort bookmarks by Category
		sort.Sort(ByCategory(bookmarks))

		for i, bookmark := range bookmarks {
			fmt.Printf("%d. [%s] %s: %s\n", i+1, bookmark.Category, bookmark.Name, bookmark.URL)
		}
	},
}

func init() {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	SitesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
