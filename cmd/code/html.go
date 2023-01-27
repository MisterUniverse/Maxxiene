package code

import (
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	html string
)

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fm.Copy("./templates/html/index.html", "../test.html")
	},
}

func init() {
	htmlCmd.Flags().StringVarP(&html, "html-basic", "b", "", "Creates a basic html boilerplate")
	// htmlCmd.Flags().StringVarP(&html, "server", "s", "", "Creates a basic tcp server")
	CodeCmd.AddCommand(htmlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// htmlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// htmlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
