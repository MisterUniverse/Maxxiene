package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Maxxiene",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Maxxiene v0.8.7")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
