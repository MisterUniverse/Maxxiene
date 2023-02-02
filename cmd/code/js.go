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
		fm.CopyFile("../../templates/js/index.html", "../../test.js")
	},
}

func init() {
	jsCmd.Flags().StringVarP(&jsc, "hello", "h", "", "Creates a basic hello world program")
	jsCmd.Flags().StringVarP(&jsc, "server", "s", "", "Creates a basic tcp server")
	CodeCmd.AddCommand(jsCmd)
}
