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
				os.WriteFile(viper.GetString("BOOKMARKS"), bookmarkJSON, 0644)
				fmt.Printf("Removed bookmark: %s\n", name)
				return
			}
		}
		fmt.Println("Bookmark not found.")
	},
}

func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData+"\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}

	SitesCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
