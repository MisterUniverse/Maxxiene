package code

import (
	fm "maxx/filemanager"

	"github.com/spf13/cobra"
)

var (
	jsc string
)

// jsCmd represents the js command
var jsCmd = &cobra.Command{
	Use:   "js",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fm.Copy("../../templates/js/index.html", "../../test.js")
	},
}

func init() {
	jsCmd.Flags().StringVarP(&jsc, "hello", "h", "", "Creates a basic hello world program")
	jsCmd.Flags().StringVarP(&jsc, "server", "s", "", "Creates a basic tcp server")
	CodeCmd.AddCommand(jsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
