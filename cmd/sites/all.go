/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package sites

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(viper.GetString("BOOKMARKS"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		json.Unmarshal(data, &bookmarks)

		for _, bookmark := range bookmarks {
			err := openBrowser(bookmark.URL)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Opened bookmark: %s\n", bookmark.Name)
		}
	},
}

// openBrowser tries to open the URL in a browser
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func init() {
	SitesCmd.AddCommand(allCmd)
}
