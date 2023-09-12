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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a bookmark",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		for i, bookmark := range bookmarks {
			if bookmark.Name == name {
				bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
				bookmarkJSON, _ := json.Marshal(bookmarks)
				os.WriteFile(viper.GetString("paths.BOOKMARKS"), bookmarkJSON, 0644)
				fmt.Printf("Removed bookmark: %s\n", name)
				return
			}
		}
		fmt.Println("Bookmark not found.")
	},
}

func init() {
	SitesCmd.AddCommand(deleteCmd)
}
