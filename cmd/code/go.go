package code

import (
	"github.com/spf13/cobra"
)

var (
	helloworld string
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Creates some boiler plate golang code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func init() {
	goCmd.Flags().StringVarP(&helloworld, "hello", "h", "", "Creates a basic hello world program")
	goCmd.Flags().StringVarP(&helloworld, "server", "s", "", "Creates a basic tcp server")
	CodeCmd.AddCommand(goCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
