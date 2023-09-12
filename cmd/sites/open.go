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

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Select a bookmark and open it in the browser",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bookmarkName := args[0]
		data, err := os.ReadFile(viper.GetString("paths.BOOKMARKS"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		json.Unmarshal(data, &bookmarks)

		for _, bookmark := range bookmarks {
			if bookmark.Name == bookmarkName {
				err := openBrowser(bookmark.URL)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				fmt.Printf("Opened bookmark: %s\n", bookmarkName)
				return
			}
		}
		fmt.Printf("No bookmark found with the name: %s\n", bookmarkName)
	},
}

func init() {
	SitesCmd.AddCommand(openCmd)
}
